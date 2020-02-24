package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Request %s method to %s\n", r.Method, r.URL.Path)
	})

	if err := http.ListenAndServe("localhost:8000", h); err != nil {

		log.Fatalf("error: listening and serving: %s", err)
	}
}
