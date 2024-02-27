/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 10:49:26
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	shellPath := "/home/xx/test.sh"
	argv := make([]string, 1)
	attr := new(os.ProcAttr)
	newProcess, err := os.StartProcess(shellPath, argv, attr) //运行脚本
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Process PID", newProcess.Pid)
	processState, err := newProcess.Wait() //等待命令执行完
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("processState PID:", processState.Pid()) //获取PID
	fmt.Println("ProcessExit:", processState.Exited())   //获取进程是否退出
}
