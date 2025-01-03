/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 15:31:52
*/
package main

import (
	"encoding/hex"
	"os"
)

func main() {
	lines := []string{
		"Go is an open source programming language.",
		"\n",
		"We encourage all Go users to subscribe to golang-announce.",
	}

	stdoutDumper := hex.Dumper(os.Stdout)

	defer stdoutDumper.Close()

	for _, line := range lines {
		stdoutDumper.Write([]byte(line))
	}

}
