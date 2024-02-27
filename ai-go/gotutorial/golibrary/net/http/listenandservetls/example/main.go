/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 10:02:34
*/
package main

import (
	"io"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

/*
生成私钥：openssl genrsa -out key.pem 2048
生成证书：openssl req -new -x509 -key key.pem -out cert.pem -days 3650
*/
func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
