/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 16:02:38
*/
package main

import (
	"fmt"
	"net/http"
	"net/http/fcgi"
	"os"
)

func myhandler(_ http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Printf("Content-type: text/html\n\n")
	fmt.Printf("<!DOCTYPE html>\n")
	fmt.Printf("<p>username: %s\n", username)
	fmt.Printf("<p>password: %s\n", password)
}

func main() {
	if err := fcgi.Serve(nil, http.HandlerFunc(myhandler)); err != nil {
		panic(err)
	}
}
