/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 19:01:15
*/
package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("sleep", "1")
	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)
}
