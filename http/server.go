package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a handler for the path "/hello"
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, X! Your Go server works 🚀")
	})

	// Start the server on port 8080
	fmt.Println("Server running on http://localhost:8080/hello")
	http.ListenAndServe(":8080", nil)
}

