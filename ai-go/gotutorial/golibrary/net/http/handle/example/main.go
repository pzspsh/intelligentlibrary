/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 16:47:49
*/
package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	params := r.URL.Query()

	// Extract specific parameters
	name := params.Get("name")
	ageParam := params.Get("age")

	// Convert age to integer
	age, err := strconv.Atoi(ageParam)
	if err != nil {
		http.Error(w, "Invalid age parameter", http.StatusBadRequest)
		return
	}

	// Process the query parameters (e.g., save them to a database, etc.)

	// Send a response to the client
	fmt.Fprintf(w, "Hello, %s! You are %d years old.", name, age)
}

func main() {
	http.HandleFunc("/parse", handleRequest)

	// Start the HTTP server
	port := ":8080"
	fmt.Printf("Starting server on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
