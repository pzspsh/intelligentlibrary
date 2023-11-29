/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 16:37:30
*/
package main

import (
	"context"
	"fmt"
)

/*
这个示例演示了如何将值传递给上下文，以及如果存在该值如何检索它。
*/
func main() {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, favContextKey("color"))

}
