/*
@File   : main.go
@Author : pan
@Time   : 2024-01-08 11:52:30
*/
package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func webServer(w http.ResponseWriter, r *http.Request) {
	timer := time.NewTimer(10 * time.Second)
	select {
	case <-r.Context().Done():
		log.Println("Error when processing request:", r.Context().Err())
		return
	case <-timer.C:
		log.Println("writing response...")
		_, err := io.WriteString(w, "Hello context")
		if err != nil {
			log.Println("Error when writing response", err)
		}
		return
	}
}

func main() {
	http.HandleFunc("/", webServer)
	log.Println("starting web server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
import (
	"io"
	"log"
	"net/http"
	"time"
)

func webServer(w http.ResponseWriter, r *http.Request) {
	timer := time.NewTimer(10 * time.Second)
	<-timer.C
	log.Println("writing response....")
	_, err := io.WriteString(w, "Hello context")
	if err != nil {
		log.Println("Error when writing response", err)
	}
}

func main() {
	http.HandleFunc("/", webServer)
	log.Println("starting web server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
*/
