/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 10:03:37
*/
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/get", getHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := r.URL.Query()
	fmt.Println(data.Get("name"))
	fmt.Println(data.Get("age"))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}
