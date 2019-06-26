package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	// If OUTBOUND is set, start a thread that creates constant outbound HTTP connections to that address
	if outbound := os.Getenv("OUTBOUND"); outbound != "" {
		go connect(outbound)
	}

	// Start a HTTP listener on port 80 that just logs all incoming requests, and their HTTP headers
	createListener()

}

// connect creates an outbound HTTP GET request every 100ms to the provided address
func connect(outbound string) {
	for {
		time.Sleep(100 * time.Millisecond)
		resp, err := http.Get(outbound)
		if err != nil {
			log.Printf("Outbound failure: %s", err)
			continue
		}
		log.Printf("Outbound connection to %s responded with HTTP %d", outbound, resp.StatusCode)
	}
}

func createListener() {

	// Create a new HTTP router
	r := mux.NewRouter()

	// Respond on "/" with a simple echo of the incoming request details
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Printf("    %s: %s\n", k, strings.Join(v, " "))
		}
	})

	// Wrap the router in a logging middleware
	logged := handlers.LoggingHandler(os.Stdout, r)

	// Start the HTTP listener
	log.Printf("Listening on port 80...")
	log.Fatal(http.ListenAndServe(":80", logged))

}
