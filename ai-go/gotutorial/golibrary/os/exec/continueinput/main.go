/*
@File   : main.go
@Author : pan
@Time   : 2024-11-20 10:54:06
*/
package main

import (
	"fmt"
	"log"
	"os/exec"
	"sync"
)

// 持续输入
func main() {
	cmd := exec.Command("openssl")
	stdout, _ := cmd.StdoutPipe() // 输出
	stderr, _ := cmd.StderrPipe() // 错误输出
	stdin, _ := cmd.StdinPipe()   // 输入
	cmd.Start()

	// 读
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		for {
			buf := make([]byte, 1024)
			n, err := stderr.Read(buf)
			if n > 0 {
				fmt.Println(string(buf[:n]))
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
		defer wg.Done()
		for {
			buf := make([]byte, 1024)
			n, err := stdout.Read(buf)
			if n == 0 {
				break
			}
			if n > 0 {
				fmt.Println(string(buf[:n]))
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

	go func() { // 写
		stdin.Write([]byte("version\n\n"))
		stdin.Write([]byte("ciphers -v\n\n"))
		stdin.Write([]byte("s_client -connect razeen.me:443"))
		stdin.Close()
		wg.Done()
	}()
	wg.Wait()
	if err := cmd.Wait(); err != nil {
		log.Printf("cmd wait %v", err)
		return
	}
}
