package main

import (
	"io"
	"net/http"
)

func forwardRequest(origin string, w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}

	forwardURL := origin + r.URL.String()
	req, err := http.NewRequest(r.Method, forwardURL, r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Copy headers from original request
	for key, values := range r.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	res, err := client.Do(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	// Read the entire body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == "GET" {
		go addToCache(r.URL.String(), CacheEntry{res.StatusCode, body, res.Header})
	}

	for key, values := range res.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.Header().Add("X-Pranay-Cache", "MISS")
	w.WriteHeader(res.StatusCode)

	w.Write(body)
}
