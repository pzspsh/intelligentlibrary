/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 16:29:22
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	const body = "Go is a general-purpose language designed with systems programming in mind."
	req, err := http.NewRequest("PUT", "http://www.example.org", strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}

	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%q", dump)

}
