/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 10:56:47
*/
package main

import (
	    "fmt"
		"io"
    "net/http"
)

func main() {
  response, err := http.Get("https://b959e645-00ae-4bc3-8a55-7224d08b1d91.mock.pstmn.io/user/1")
    if err != nil {
        fmt.Println(err)
        panic(err)
    }
    defer response.Body.Close()
	fmt.Println(response)
    fmt.Println(response.StatusCode)
    fmt.Println(response.Status)
    fmt.Println(response.Header)
    body, err := io.ReadAll(response.Body)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(body))
}
