package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.Count([]byte("hi go go go go go go go go go"), []byte("go"))) // 9
}
