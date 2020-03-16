package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
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
	fmt.Fprintf(os.Stdout, "Starting listening on Port 8080\n")
	http.Handle("/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handler)))
	http.HandleFunc("/health", healthcheckHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
