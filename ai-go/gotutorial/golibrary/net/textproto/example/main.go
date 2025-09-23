package main

import (
	"bufio"
	"fmt"
	"net/textproto"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string with leading and trailing whitespace: ")

	input, _ := reader.ReadString('\n')
	trimmedBytes := textproto.TrimBytes([]byte(input))
	trimmedString := string(trimmedBytes)
	fmt.Printf("Trimmed string: '%s'\n", trimmedString)
}
