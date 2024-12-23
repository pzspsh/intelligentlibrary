/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 16:56:32
*/
package main

import (
	"fmt"
	"net/http"
)

func customStatusHandler(w http.ResponseWriter, r *http.Request) {
	// Define a custom status code
	customStatusCode := 418

	// Retrieve the status text for the custom status code
	statusText := http.StatusText(customStatusCode)

	// Set the status code and status text in the response
	w.WriteHeader(customStatusCode)
	fmt.Fprintf(w, "Status code: %d\nStatus text: %s", customStatusCode, statusText)
}

func main() {
	// Register the customStatusHandler function for the "/custom-status" route
	http.HandleFunc("/custom-status", customStatusHandler)

	// Start the web server
	fmt.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", nil)
}
