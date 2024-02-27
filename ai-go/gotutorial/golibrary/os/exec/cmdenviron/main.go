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
	cmd := exec.Command("pwd")

	// Set Dir before calling cmd.Environ so that it will include an
	// updated PWD variable (on platforms where that is used).
	cmd.Dir = ".."
	cmd.Env = append(cmd.Environ(), "POSIXLY_CORRECT=1")

	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)
}
