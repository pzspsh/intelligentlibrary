package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.Compare([]byte{}, []byte{}))   // 0
	fmt.Println(bytes.Compare([]byte{1}, []byte{2})) // -1
	fmt.Println(bytes.Compare([]byte{2}, []byte{1})) // 1
	fmt.Println(bytes.Compare([]byte{}, nil))        //0
}
