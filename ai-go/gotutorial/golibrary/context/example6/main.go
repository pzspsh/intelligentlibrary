/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 17:39:27
*/
package main

import (
	"context"

	"fmt"

	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)

	defer cancel()

	resultChan := make(chan string)

	go func() {

		operation(ctx, resultChan)

	}()

	select {

	case res := <-resultChan:

		fmt.Println("Operation completed:", res)

	case <-ctx.Done():

		fmt.Println("Operation timed out:", ctx.Err())

	}

}

func operation(ctx context.Context, resultChan chan string) {

	select {

	case <-time.After(200 * time.Millisecond):

		resultChan <- "Success"

	case <-ctx.Done():

		return

	}

}
