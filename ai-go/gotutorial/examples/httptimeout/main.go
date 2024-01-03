/*
@File   : main.go
@Author : pan
@Time   : 2024-01-03 12:10:16
*/
package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
	// "io/ioutil"
)

func main() {
	/*
	   c := make(chan struct{})
	   	timer := time.AfterFunc(5*time.Second, func() {
	   		close(c)
	   	})

	   	// Serve 256 bytes every second.
	   	req, err := http.NewRequest("GET", "http://httpbin.org/range/2048?duration=8&chunk_size=256", nil)
	   	if err != nil {
	   		log.Fatal(err)
	   	}
	   	if err != nil {
	   		fmt.Println(err)
	   	}
	   	req.Cancel = c
	*/
	ctx, cancel := context.WithCancel(context.TODO())
	timer := time.AfterFunc(5*time.Second, func() {
		cancel()
	})
	req, err := http.NewRequest("GET", "http://httpbin.org/range/2048?duration=8&chunk_size=256", nil)
	if err != nil {
		log.Fatal(err)
	}
	req = req.WithContext(ctx)
	log.Println("Sending request...")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	log.Println("Reading body...")
	for {
		timer.Reset(2 * time.Second)
		// Try instead: timer.Reset(50 * time.Millisecond)
		// _, err = io.CopyN(ioutil.Discard, resp.Body, 256)
		_, err = io.CopyN(io.Discard, resp.Body, 256)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
}
