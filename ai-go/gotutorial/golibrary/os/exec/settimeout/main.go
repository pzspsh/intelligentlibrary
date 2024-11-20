/*
@File   : main.go
@Author : pan
@Time   : 2024-11-20 10:52:22
*/
package main

import (
	"bufio"
	"context"
	"log"
	"os/exec"
	"time"
)

// 通过上下文控制超时
func main() {
	ctx, calcel := context.WithTimeout(context.Background(), 2*time.Second)
	defer calcel()
	cmd := exec.CommandContext(ctx, "./testcmd/testcmd", "-s", "-e")
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	oReader := bufio.NewReader(stdout)
	eReader := bufio.NewReader(stderr)
	cmd.Start()
	go func() {
		for {
			line, err := oReader.ReadString('\n')
			if line != "" {
				log.Printf("read line %s", line)
			}
			if err != nil || line == "" {
				log.Printf("read line err %v", err)
				return
			}

		}
	}()

	go func() {
		for {
			line, err := eReader.ReadString('\n')
			if line != "" {
				log.Printf("read err %s", line)
			}
			if err != nil || line == "" {
				log.Printf("read err %v", err)
				return
			}

		}
	}()
	if err := cmd.Wait(); err != nil {
		log.Printf("cmd wait %v", err)
		return
	}
}
