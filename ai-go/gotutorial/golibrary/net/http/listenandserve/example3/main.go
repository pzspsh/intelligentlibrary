/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 16:46:07
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Golang HTTP Server!")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
