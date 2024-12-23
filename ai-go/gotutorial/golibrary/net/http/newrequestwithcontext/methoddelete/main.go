/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 12:25:56
*/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("4. Performing Http Delete...")
	todo := Todo{1, 2, "lorem ipsum dolor sit amet", true}
	jsonReq, err := json.Marshal(todo)
	if err != nil {
		fmt.Println(err)
	}
	req, err := http.NewRequest(http.MethodDelete, "https://jsonplaceholder.typicode.com/todos/1", bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Println(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}
