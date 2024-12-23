/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 15:31:52
*/
package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	const s = "48656c6c6f20476f7068657221"
	decoded, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", decoded)

}
