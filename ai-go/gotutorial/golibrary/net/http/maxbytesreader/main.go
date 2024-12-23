/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 13:55:05
*/
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/handle", maxBytes(PostHandler)).Methods("POST")
	http.ListenAndServe(":8080", r)
}

// Middleware to enforce the maximum post body size

// func maxBytes(f http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// As an example, limit post body to 10 bytes
// 		r.Body = http.MaxBytesReader(w, r.Body, 10)
// 		f(w, r)
// 	}
// }

func maxBytes(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 10)
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		f(w, r)
	}
}

// func PostHandler(w http.ResponseWriter, r *http.Request) {
// 	// How do I know if the form data has been truncated?
// 	book := r.FormValue("email")
// 	fmt.Fprintf(w, "You've requested the book: %s\n", book)
// }

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	book := r.FormValue("email")
	fmt.Fprintf(w, "You've requested the book: %s\n", book)
}
