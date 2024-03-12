/*
@File   : main.go
@Author : pan
@Time   : 2024-03-12 15:53:21
*/
package main

import (
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
