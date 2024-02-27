/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 11:17:21
*/
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	client := http.Client{}

	user := User{
		Username: "123456",
		Password: "12346",
	}
	dataByte, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	bodyReader := bytes.NewReader(dataByte)
	// http.MethodConnect
	// http.MethodDelete
	// http.MethodGet
	// http.MethodHead
	// http.MethodOptions
	// http.MethodPatch
	// http.MethodPost
	// http.MethodPut
	// http.MethodTrace
	request, err := http.NewRequestWithContext(context.Background(), http.MethodPost, "http://localhost:8080/user", bodyReader)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("statusCode: ", resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	fmt.Println(string(body))
}
