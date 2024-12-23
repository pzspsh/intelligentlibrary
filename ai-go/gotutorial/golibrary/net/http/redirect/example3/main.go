/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 16:38:40
*/
package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", redirectHandler)
	http.ListenAndServe(":8080", nil)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	targetURL := "https://www.example.com"
	http.Redirect(w, r, targetURL, http.StatusFound)
}
