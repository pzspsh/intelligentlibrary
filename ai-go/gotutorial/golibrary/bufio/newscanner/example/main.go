/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 17:07:33
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if ok := scanner.Scan(); !ok {
			break
		}
		str := scanner.Text()
		reader := strings.NewReader(str)
		for {
			runeScanner := bufio.NewReader(reader)
			r, _, err := runeScanner.ReadRune()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}
			fmt.Printf("%#U\n", r)
		}
	}
}
