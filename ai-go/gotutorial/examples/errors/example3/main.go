/*
@File   : main.go
@Author : pan
@Time   : 2024-01-08 11:59:52
*/
package main

import (
	"errors"
	"fmt"
)

type DBError struct {
	desc string
}

func (dbError DBError) Error() string {
	return dbError.desc
}

func getRecord() error {
	return DBError{
		desc: "no rows found",
	}
}

func webService() error {
	if err := getRecord(); err != nil {
		return fmt.Errorf("error %w when calling DB", err)
	}
	return nil
}

func main() {
	if err := webService(); err != nil {
		var dbError DBError
		if errors.As(err, &dbError) {
			fmt.Printf("The searched record cannot be found. Error returned from DB is %s", dbError)
			return
		}
		fmt.Println("unknown error when searching record")
		return

	}
	fmt.Println("webservice call successful")

}
