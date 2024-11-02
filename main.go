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
			fmt.Printf("GET request from %s\n", r.URL.Path)
			// http.Get()
		} else {
			fmt.Printf("%s request from %s\n", r.Method, r.URL.Path)
		}
	})

	fmt.Printf("Hello! The server is running on port %d.\n", *port)
	fmt.Printf("All GET request caching from %s starting now...\n", *origin)

	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
