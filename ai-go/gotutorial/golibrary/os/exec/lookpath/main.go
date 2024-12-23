/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 19:01:15
*/
package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	path, err := exec.LookPath("fortune")
	if err != nil {
		log.Fatal("installing fortune is in your future")
	}
	fmt.Printf("fortune is available at %s\n", path)
}
