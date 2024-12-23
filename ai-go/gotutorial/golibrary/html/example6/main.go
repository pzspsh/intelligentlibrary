/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 10:48:29
*/
package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type Teacher struct {
	Name    string
	Subject string
}
type Student struct {
	Id      int
	Name    string
	Country string
}

type Rooster struct {
	Teacher  Teacher
	Students []Student
}

func ShowIndexView(response http.ResponseWriter, request *http.Request) {
	teacher := Teacher{
		Name:    "Alex",
		Subject: "Physics",
	}
	students := []Student{
		{Id: 1001, Name: "Peter", Country: "China"},
		{Id: 1002, Name: "Jeniffer", Country: "Sweden"},
	}
	rooster := Rooster{
		Teacher:  teacher,
		Students: students,
	}

	tmpl, err := template.ParseFiles("./views/layout.html",
		"./views/nav.html",
		"./views/content.html")

	if err != nil {
		fmt.Println("Error " + err.Error())
	}
	tmpl.Execute(response, rooster)
}

func RegisterRoutes(r *mux.Router) {
	viewRouter := r.PathPrefix("/view").Subrouter()
	viewRouter.HandleFunc("/index", ShowIndexView)
}

func main() {

}
