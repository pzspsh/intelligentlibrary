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
	const text = "<<.Greeting>> {{.Name}}"

	data := struct {
		Greeting string
		Name     string
	}{
		Greeting: "Hello",
		Name:     "Joe",
	}

	t := template.Must(template.New("tpl").Delims("<<", ">>").Parse(text))

	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}

}
