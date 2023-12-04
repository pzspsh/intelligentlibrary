/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 16:54:39
*/
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

/*
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365
*/
func main() {
	// Define a simple handler function
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	// Create a new http.Server with custom configurations
	server := &http.Server{
		Addr:    ":8443",
		Handler: handler,
	}

	// Prepare the TLS certificate and private key file paths
	certFile := "./cert.pem"
	keyFile := "./key.pem"

	// Create a listener for incoming connections
	listener, err := net.Listen("tcp", server.Addr)
	if err != nil {
		log.Fatalf("Failed to create listener: %v", err)
	}

	// Start the HTTPS server using http.ServeTLS
	fmt.Println("Server running on https://localhost:8443")
	err = server.ServeTLS(listener, certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
