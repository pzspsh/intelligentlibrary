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
