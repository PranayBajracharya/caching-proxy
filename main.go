package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "The port to listen on")
	origin := flag.String("origin", "", "The origin to cache GET requests from")

	flag.Parse()

	if *origin == "" {
		fmt.Fprintln(os.Stderr, "Error: --origin is required")
		flag.Usage()
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			fmt.Printf("GET request from %s\n", r.URL.Path)
		} else {
			forwardRequest(*origin, w, r)
			fmt.Printf("%s request from %s\n", r.Method, r.URL.Path)
			return
		}
	})

	fmt.Printf("Hello! The server is running on port %d.\n", *port)
	fmt.Printf("All GET request caching from %s starting now...\n", *origin)

	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func forwardRequest(origin string, w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}

	forwardURL := origin + r.URL.Path
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

	w.WriteHeader(res.StatusCode)
	for key, values := range res.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	io.Copy(w, res.Body)
}
