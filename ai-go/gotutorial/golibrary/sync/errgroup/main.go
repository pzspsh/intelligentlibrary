package main

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	var g errgroup.Group

	urls := []string{"https://example.com", "https://example.org", "https://example.net"}

	for _, url := range urls {
		// 注意重新声明变量，避免闭包问题
		g.Go(func() error {
			// 模拟一个可能失败的操作，例如HTTP请求
			resp, err := http.Get(url)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			fmt.Printf("Successfully fetched %s\n", url)
			return nil
		})
	}

	// Wait will return the first non-nil error returned by a goroutine.
	// If all goroutines succeed, Wait returns nil.
	if err := g.Wait(); err == nil {
		fmt.Println("All requests succeeded!")
	} else {
		fmt.Printf("There was an error: %v\n", err)
	}
}
