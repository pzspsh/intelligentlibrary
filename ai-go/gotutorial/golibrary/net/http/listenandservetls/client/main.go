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

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Golang HTTPS Server!")
	})

	err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
