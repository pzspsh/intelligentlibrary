/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 16:31:34
*/
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	h := json.RawMessage(`{"precomputed": true}`)

	c := struct {
		Header *json.RawMessage `json:"header"`
		Body   string           `json:"body"`
	}{Header: &h, Body: "Hello Gophers!"}

	b, err := json.MarshalIndent(&c, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)

}
