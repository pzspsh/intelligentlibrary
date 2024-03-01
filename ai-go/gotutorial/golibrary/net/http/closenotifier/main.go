/*
@File   : main.go
@Author : pan
@Time   : 2024-03-01 10:57:09
*/
package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Content-Type-Options", "nosniff")

	fmt.Println("Client connected!")

	flusher, ok := w.(http.Flusher)
	if !ok {
		fmt.Println("ResponseWriter doesn't implement Flusher interface")
		return
	}

	closeNotifier, ok := w.(http.CloseNotifier)
	if !ok {
		fmt.Println("ResponseWriter doesn't implement CloseNotifier interface")
		return
	}
	closeNotifyChannel := closeNotifier.CloseNotify()

	for {
		fmt.Println("Sending data chunk...")
		fmt.Fprintf(w, "Chunk.")
		flusher.Flush()

		select {
		case <-closeNotifyChannel:
			goto closed
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}

closed:
	fmt.Println("Client disconnected")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	server := &http.Server{
		Addr:    "localhost:8000",
		Handler: mux,
		ConnState: func(conn net.Conn, state http.ConnState) {
			fmt.Printf("[ConnState] %v: ", conn.RemoteAddr())

			switch state {
			case http.StateNew:
				fmt.Println("StateNew")
			case http.StateActive:
				fmt.Println("StateActive")
			case http.StateIdle:
				fmt.Println("StateIdle")
			case http.StateHijacked:
				fmt.Println("StateHijacked")
			case http.StateClosed:
				fmt.Println("StateClosed")
			}
		},
	}

	server.ListenAndServe()
}
