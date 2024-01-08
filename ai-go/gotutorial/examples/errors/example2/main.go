/*
@File   : main.go
@Author : pan
@Time   : 2024-01-08 11:58:52
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
		return fmt.Errorf("error %w when calling DB", err)
	}
	return nil
}

func main() {
	if err := webService(); err != nil {
		if errors.Is(err, errFoo) {
			fmt.Printf("The searched record cannot be found. Error returned from DB is %s", err)
			return
		}
		fmt.Println("unknown error when searching record")
		return

	}
	fmt.Println("webservice call successful")
}
