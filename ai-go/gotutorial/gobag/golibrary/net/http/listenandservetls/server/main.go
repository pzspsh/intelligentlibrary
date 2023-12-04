/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 16:45:13
*/
package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	server := &http.Server{
		Addr:    ":8443",
		Handler: mux,
	}

	fmt.Println("Starting HTTPS server on https://localhost:8443")
	err := server.ListenAndServeTLS("cert.pem", "key.pem")

	if err != nil && err != http.ErrServerClosed {
		fmt.Printf("Error starting HTTPS server: %v\n", err)
	}
}
