package main

import (
	"fmt"
	"net/http"
)

func route(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Vercel!")
}

func main() {
	http.HandleFunc("/api/hello", route)
	http.ListenAndServe(":8080", nil)
}
