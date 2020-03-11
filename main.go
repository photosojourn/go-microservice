package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if os.Getenv("GO_MS_MESSAGE") == "" {
		fmt.Fprintf(w, "Hi there this is a microservice in Golang\n")
	} else {
		fmt.Fprintf(w, os.Getenv("GO_MS_MESSAGE"))
	}
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/health", healthcheckHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
