package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./test.tmpl")
	if err != nil {
		fmt.Printf("parse file failed err := %v", err)
	}
	mp := map[string]any{
		"name": "张三",
		"sex":  "男",
	}
	err = t.Execute(w, mp)
	if err != nil {
		fmt.Printf("execute file failed err := %v", err)
	}

}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("listen address failed err = %v", err)
	}
}
