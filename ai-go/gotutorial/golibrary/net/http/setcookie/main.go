/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 16:55:46
*/
package main

import (
	"fmt"
	"net/http"
	"time"
)

func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new cookie
	cookie := &http.Cookie{
		Name:     "welcomeMessage",
		Value:    "Hello, Gopher!",
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	}

	// Set the cookie using the http.SetCookie function
	http.SetCookie(w, cookie)

	// Respond to the client
	fmt.Fprint(w, "Cookie has been set!")
}

func main() {
	// Register the setCookieHandler function for the "/set-cookie" route
	http.HandleFunc("/set-cookie", setCookieHandler)

	// Start the web server
	fmt.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", nil)
}
