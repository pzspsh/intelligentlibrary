/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 16:55:06
*/
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	if _, err := os.Open("non-existing"); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("file does not exist")
		} else {
			fmt.Println(err)
		}
	}

}
