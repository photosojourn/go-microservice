package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there this is a microservice in Golang\n")
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/health", healthcheckHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
