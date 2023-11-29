/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 17:22:10
*/
package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func HttpRequest(
	ctx context.Context,
	req *http.Request,
	client *http.Client,
	respCh chan []byte,
	errCh chan error,
) {
	req = req.WithContext(ctx)
	tr := &http.Transport{}
	client.Transport = tr

	go func() {
		resp, err := client.Do(req)
		if err != nil {
			log.Println("http.Client.Do failure, err:", err)
			errCh <- err
		}
		if resp != nil {
			defer resp.Body.Close()
			respData, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Println("resp read all failure.err:", err)
				errCh <- err
			}
			respCh <- respData
		} else {
			errCh <- errors.New("http request failure")
		}
	}()
	for {
		select {
		case <-ctx.Done():
			tr.CancelRequest(req)
			errCh <- errors.New("http request Canceled")
			return
		case <-errCh:
			fmt.Println("into cancel request")
			tr.CancelRequest(req)
			return
		}
	}
}

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://127.0.0.1:8080/get/user", nil)
	if err != nil {
		log.Println("http new request failue ,err:", err)
	}
	ctx, _ := context.WithCancel(context.Background())
	respCh := make(chan []byte)
	errCh := make(chan error)
	go func() {
		select {
		case <-errCh:
			fmt.Println("into there")
			//cancel()
		case resp := <-respCh:
			fmt.Println("resp:", string(resp))
		}
	}()
	HttpRequest(ctx, req, client, respCh, errCh)
}
