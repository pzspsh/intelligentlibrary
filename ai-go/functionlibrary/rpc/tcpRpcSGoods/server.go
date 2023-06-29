/*
@File   : server.go
@Author : pan
@Time   : 2023-06-29 14:51:18
*/
package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// goods微服务:服务端,传入struct,增加商品,获取商品

// 创建远程调用的函数，函数一般是放在结构体里面
type Goods struct{}

// AddGoods参数对应的结构体
// 增加商品请求参数结构体
type AddGoodsReq struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

// 增加商品返回结构体
type AddGoodsRes struct {
	Success bool
	Message string
}

// GetGoods参数对应的结构体
// 获取商品请求结构体
type GetGoodsReq struct {
	Id int
}

// 获取商品返回结构体
type GetGoodsRes struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

/*
说明:
    1、方法只能有两个可序列化的参数，其中第二个参数是指针类型
        req 表示获取客户端传过来的数据
        res 表示给客户端返回数据
    2、方法要返回一个error类型，同时必须是公开的方法
    3、req和res的类型不能是：channel（通道）、func（函数）,因为以上类型均不能进行 序列化
*/

// 增加商品函数
func (g Goods) AddGoods(req AddGoodsReq, res *AddGoodsRes) error {
	//1、执行增加 模拟
	fmt.Printf("%#v\n", req)
	*res = AddGoodsRes{
		Success: true, //根据增加结果,返回状态
		Message: "增加商品成功",
	}
	return nil
}

// 获取商品函数
func (g Goods) GetGoods(req GetGoodsReq, res *GetGoodsRes) error {
	//1、执行获取商品 模拟
	fmt.Printf("%#v\n", req)

	//2、返回获取的结果
	*res = GetGoodsRes{
		Id:      12, //商品id
		Title:   "服务器获取的数据",
		Price:   24.5,
		Content: "我是服务器数据库获取的内容",
	}
	return nil
}

func main() {
	//1、 注册RPC服务
	//goods: rpc服务名称
	err := rpc.RegisterName("goods", new(Goods))
	if err != nil {
		fmt.Println(err)
	}

	//2、监听端口
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
	}
	//3、应用退出的时候关闭监听端口
	defer listen.Close()

	for { // for 循环, 一直进行连接,每个客户端都可以连接
		fmt.Println("准备建立连接")
		//4、建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
		}
		//5、绑定服务
		rpc.ServeConn(conn)
	}
}
