package main

import (
	"fmt"
	"go/importer"
)

func main() {
	pkg, err := importer.Default().Import("github.com/onsi/ginkgo")
	if err != nil {
		panic(err)
	}
	fmt.Println(pkg)
}
