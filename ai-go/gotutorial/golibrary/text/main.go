/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 15:20:28
*/
package main

import (
	"os"
	"text/template"
)

func main() {
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// Must是上面包装err成panic的简便写法
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{"Go", "Rust", "C++", "C#"})

	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// 支持用字段名替换
	t2 := Create("t2", "Name: {{.Name}}\n")
	t2.Execute(os.Stdout, struct{ Name string }{"Jane Doe"})
	t2.Execute(os.Stdout, map[string]string{"Name": "Mickey Mouse"})

	// 支持基于条件判断的替换
	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	// 支持基于遍历的替换
	t4 := Create("t4", "Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout, []string{"Go", "Rust", "C++", "C#"})
}
