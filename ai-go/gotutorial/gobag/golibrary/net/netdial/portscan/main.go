/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 01:36:18
*/
package main

/* import (
	"fmt"
	"net"
)

func main() {
	//通过conn和err读取net.Dial()函数的返回值
	conn, err := net.Dial("tcp", "scanme.nmap.org:80")

	//通过err的值检验端口是否连接成功
	if err == nil {
		fmt.Println("Connection successful")

		//如果端口连接成功，则需要主动关闭端口连接
		conn.Close()
	}
}
*/

/* import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		//将i的值和地址一起转换为字符串
		address := fmt.Sprintf("scanme.nmap.org:%d", i)

		conn, err := net.Dial("tcp", address)

		//如果端口已关闭或已过滤，直接进入下一循环
		if err != nil {
			continue
		}

		//关闭端口连接
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
} */

import (
	"fmt"
	"net"
	"time"
)

func main() {
	for i := 1; i <= 1024; i++ {

		go func(j int) {
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)

			//如果端口已关闭或已过滤，结束本次循环
			if err != nil {
				return
			}

			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)

	}
	//强制函数main()暂停1s，以保证匿名函数可以执行完成
	time.Sleep(1 * time.Second)
}
