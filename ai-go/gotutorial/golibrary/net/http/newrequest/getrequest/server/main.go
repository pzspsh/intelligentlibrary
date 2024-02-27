/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 10:06:17
*/
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/get/", getHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fmt.Println(r.URL)

	//获取查询参数
	data := r.URL.Query()
	fmt.Println(data.Get("username"))
	fmt.Println(data.Get("password"))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}
