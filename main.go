package main

import (
	"flag"
	"fmt"
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
			entry, ok := getFromCache(r.URL.String())
			if ok {
				sendCachedResponse(w, entry)
				return
			}
			forwardRequest(*origin, w, r)
			fmt.Printf("GET request from %s\n", r.URL.String())
		} else {
			forwardRequest(*origin, w, r)
			fmt.Printf("%s request from %s\n", r.Method, r.URL.String())
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
