package main

import (
	"io"
	"log"
	"net/http"
)

const VERSION string = "1.0.6"

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Current version: "+VERSION)
}

func main() {
	log.Printf("Listening on port 8000...")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
