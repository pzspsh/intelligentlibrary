/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 12:34:59
*/
package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func main() {
	// 1.
	payload, err := json.Marshal(map[string]interface{}{
		"title":     "my simple todo",
		"completed": false,
	})
	if err != nil {
		log.Fatal(err)
	}

	// 2.
	client := &http.Client{}
	url := "https://sample42.free.beeceptor.com/trace"

	// 3.
	req, err := http.NewRequest(http.MethodTrace, url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	// 4.
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// 5.
	defer resp.Body.Close()

	// 6.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}
