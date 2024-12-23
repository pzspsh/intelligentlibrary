/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 12:27:47
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
	fmt.Println("3. Performing Http Put...")
	todo := Todo{1, 2, "lorem ipsum dolor sit amet", true}
	jsonReq, err := json.Marshal(todo)
	if err != nil {
		fmt.Println(err)
	}
	req, err := http.NewRequest(http.MethodPut, "https://jsonplaceholder.typicode.com/todos/1", bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
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

	// Convert response body to Todo struct
	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("API Response as struct:\n%+v\n", todoStruct)
}
