/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 19:01:15
*/
package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	if err := exec.CommandContext(ctx, "sleep", "5").Run(); err != nil {
		fmt.Println(err)
		// This will fail after 100 milliseconds. The 5 second sleep
		// will be interrupted.
	}
}
