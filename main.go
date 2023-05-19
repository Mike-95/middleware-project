package main

import (
	"fmt"
	"github.com/Mike-95/middleware-project/middleware"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	logger := middleware.Logger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	}))

	logger.ServeHTTP(w, r)
}
