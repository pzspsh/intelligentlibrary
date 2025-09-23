package main

import (
	"log"
	"os"
)

func main() {
	err := os.WriteFile("test.txt", []byte("Hi\n"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
