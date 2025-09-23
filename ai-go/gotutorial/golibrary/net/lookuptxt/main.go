package main

import (
	"fmt"
	"net"
)

func main() {
	txts, err := net.LookupTXT("google.com")
	if err != nil {
		panic(err)
	}
	if len(txts) == 0 {
		fmt.Printf("no record")
	}
	for _, txt := range txts {
		fmt.Printf("%s\n", txt)
	}
}
