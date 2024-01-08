/*
@File   : main.go
@Author : pan
@Time   : 2024-01-08 11:56:47
*/
package main

import (
	"errors"
	"fmt"
)

var errFoo = errors.New("no rows found")

func getRecord() error {
	return errFoo
}

func webService() error {
	if err := getRecord(); err != nil {
		return fmt.Errorf("error %s when calling DB", err)
	}
	return nil
}

func main() {
	if err := webService(); err != nil {
		fmt.Printf("Error: %s when calling webservice\n", err)
		return
	}
	fmt.Println("webservice call successful")
}
