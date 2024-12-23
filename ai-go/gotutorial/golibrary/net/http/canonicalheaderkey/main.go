/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:05:19
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	hea1 := http.CanonicalHeaderKey("uid-test")
	hea2 := http.CanonicalHeaderKey("accept-encoding")
	fmt.Println(hea1) //Uid-Test
	fmt.Println(hea2) //Accept-Encoding
}
