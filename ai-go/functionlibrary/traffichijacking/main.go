/*
@File   : main.go
@Author : pan
@Time   : 2024-04-29 11:02:12
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("BBBBBBBBB", r.URL.Path)
	fmt.Println(r.Header)
	// fmt.Println(resp.Header)
	// Connect to Burp Suite proxy
	proxyConn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("Error connecting to Burp Suite proxy: %v", err)
	}
	defer proxyConn.Close()

	// Forward request to Burp Suite
	err = r.WriteProxy(proxyConn)
	if err != nil {
		log.Fatalf("Error forwarding request to Burp Suite: %v", err)
	}

	// Read response from Burp Suite
	resp, err := http.ReadResponse(bufio.NewReader(proxyConn), r)
	if err != nil {
		log.Fatalf("Error reading response from Burp Suite: %v", err)
	}

	// Forward response to client
	for k, v := range resp.Header {
		w.Header().Set(k, v[0])
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
	resp.Body.Close()
}

func main() {
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
