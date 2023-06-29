/*
@File   : client.go
@Author : pan
@Time   : 2023-06-29 14:43:55
*/
package main

import (
	"fmt"
	"net/rpc"
)

// rpc服务端
func main() {
	//1、用 rpc.Dial和rpc微服务端建立连接
	conn, err1 := rpc.Dial("tcp", "127.0.0.1:8080")
	if err1 != nil {
		fmt.Println(err1)
	}
	//2、当客户端退出的时候关闭连接
	defer conn.Close()

	//3、调用远程函数
	//微服务端返回的数据
	var reply string
	/*
	   1、第一个参数: hello.SayHello,hello 表示服务名称  SayHello 方法名称
	   2、第二个参数: 给服务端的req传递数据
	   3、第三个参数: 需要传入地址,获取微服务端返回的数据
	*/
	err2 := conn.Call("hello.SayHello", "我是客户端", &reply)
	if err2 != nil {
		fmt.Println(err2)
	}
	//4、获取微服务返回的数据
	fmt.Println(reply)
}
