package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Open a file
	file, err := os.Open("example.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Get the file's content type
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	contentType := http.DetectContentType(buffer)

	fmt.Printf("Content Type: %s\n", contentType)
}
