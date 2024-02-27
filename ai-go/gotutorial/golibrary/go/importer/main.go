/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 17:51:36
*/
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
