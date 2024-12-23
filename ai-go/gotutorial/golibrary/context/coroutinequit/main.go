/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 16:58:06
*/
package main

import (
	"context"
	"errors"
	"sync"
)

func request(ctx context.Context, url string) error {
	result := make(chan int)
	err := make(chan error)
	go func() {
		// 假如isSuccess是请求返回的结果，成功则通过result传递成功信息，错误通过error传递错误信息
		isSuccess := true
		if isSuccess {
			result <- 1
		} else {
			err <- errors.New("some error happen")
		}
	}()

	select {
	case <-ctx.Done():
		// 其他请求失败
		return ctx.Err()
	case e := <-err:
		// 本次请求失败，返回错误信息
		return e
	case <-result:
		// 本此请求成功，不返回错误信息
		return nil
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// 调用接口a
	err := request(ctx, "https://xxx.com/a")
	if err != nil {
		return
	}
	wg := sync.WaitGroup{}
	// 调用接口b
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := request(ctx, "https://xxx.com/b")
		if err != nil {
			cancel()
		}
	}()
	// 调用接口c
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := request(ctx, "https://xxx.com/c")
		if err != nil {
			cancel()
		}
	}()
	wg.Wait()
}
