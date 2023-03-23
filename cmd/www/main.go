package main

import (
	"log"
	"net/http"

	"github.com/bmizerany/pat"
)

func main() {
	// Initialize a router and add the path and handler for the homepage.
	mux := pat.New()
	mux.Get("/:locale", http.HandlerFunc(handleHome))

	// Start the HTTP server using the router
	log.Print("starting server on :4018...")
	err := http.ListenAndServe(":4018", mux)
	log.Fatal(err)
}
