/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 10:29:57
*/
package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

// templateFile defines the contents of a template to be stored in a file, for testing.
type templateFile struct {
	name     string
	contents string
}

func createTestDir(files []templateFile) string {
	dir, err := os.MkdirTemp("", "template")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		f, err := os.Create(filepath.Join(dir, file.name))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		_, err = io.WriteString(f, file.contents)
		if err != nil {
			log.Fatal(err)
		}
	}
	return dir
}

func main() {
	// Here we create different temporary directories and populate them with our sample
	// template definition files; usually the template files would already
	// exist in some location known to the program.
	dir1 := createTestDir([]templateFile{
		// T1.tmpl is a plain template file that just invokes T2.
		{"T1.tmpl", `T1 invokes T2: ({{template "T2"}})`},
	})

	dir2 := createTestDir([]templateFile{
		// T2.tmpl defines a template T2.
		{"T2.tmpl", `{{define "T2"}}This is T2{{end}}`},
	})

	// Clean up after the test; another quirk of running as an example.
	defer func(dirs ...string) {
		for _, dir := range dirs {
			os.RemoveAll(dir)
		}
	}(dir1, dir2)

	// Here starts the example proper.
	// Let's just parse only dir1/T0 and dir2/T2
	paths := []string{
		filepath.Join(dir1, "T1.tmpl"),
		filepath.Join(dir2, "T2.tmpl"),
	}
	tmpl := template.Must(template.ParseFiles(paths...))

	err := tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
}
