/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 16:58:45
*/
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/index", indexHandle)

	log.Println("HTTPS server up at https://localhost:8443")
	http.ListenAndServeTLS(":8443", "certpath/server.crt", "certpath/server.key", nil)
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr, r.URL, r.Proto, r.Header, r.Body)
	fmt.Fprint(w, "Hi, This is an example of https service in golang!")
}
