/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 10:29:57
*/
package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	check(err)
	err = t.ExecuteTemplate(os.Stdout, "T", "<script>alert('you have been pwned')</script>")
	check(err)
}
