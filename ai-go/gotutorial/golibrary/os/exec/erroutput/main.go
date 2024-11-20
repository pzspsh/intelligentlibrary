/*
@File   : main.go
@Author : pan
@Time   : 2024-11-20 10:49:48
*/
package main

import (
	"log"
	"os/exec"
)

// stdout & stderr 分开输出
// 分离标准输出与错误输出
func main() {
	cmd := exec.Command("./testcmd/testcmd", "-s", "-e")
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	cmd.Start()

	go func() {
		for {
			buf := make([]byte, 1024)
			n, err := stderr.Read(buf)
			if n > 0 {
				log.Printf("read err %s", string(buf[:n]))
			}
			if n == 0 {
				break
			}
			if err != nil {
				log.Printf("read err %v", err)
				return
			}
		}
	}()

	go func() {
		for {
			buf := make([]byte, 1024)
			n, err := stdout.Read(buf)
			if n == 0 {
				break
			}
			if n > 0 {
				log.Printf("read out %s", string(buf[:n]))

			}
			if n == 0 {
				break
			}
			if err != nil {
				log.Printf("read out %v", err)
				return
			}

		}
	}()
	if err := cmd.Wait(); err != nil {
		log.Printf("cmd wait %v", err)
		return
	}
}
