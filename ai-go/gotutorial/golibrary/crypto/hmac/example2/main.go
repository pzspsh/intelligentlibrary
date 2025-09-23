package main

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	key := "kuteng"
	data := "www.5lmh.com"
	hmac := hmac.New(md5.New, []byte(key))
	hmac.Write([]byte(data))
	fmt.Println(hex.EncodeToString(hmac.Sum([]byte(""))))
}
