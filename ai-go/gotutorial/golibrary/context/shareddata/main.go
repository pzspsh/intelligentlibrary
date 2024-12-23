/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 17:19:19
*/
package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	process(ctx)

	ctx = context.WithValue(ctx, "traceID", "foo")
	process(ctx)
}

func process(ctx context.Context) {
	traceId, ok := ctx.Value("traceID").(string)
	if ok {
		fmt.Printf("process over. trace_id=%s\n", traceId)
	} else {
		fmt.Printf("process over. no trace_id\n")
	}
}
