/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 17:09:59
*/
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fs := flag.NewFlagSet("ExampleBoolFunc", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)

	fs.BoolFunc("log", "logs a dummy message", func(s string) error {
		fmt.Println("dummy message:", s)
		return nil
	})
	fs.Parse([]string{"-log"})
	fs.Parse([]string{"-log=0"})

}
