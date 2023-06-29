/*
@File   : client.go
@Author : pan
@Time   : 2023-06-29 14:50:50
*/
package main

import (
	"fmt"
	"net/rpc"
)

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

func main() {
	//1、用 rpc.Dial和rpc微服务端建立连接
	conn, err := rpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
	}
	//2、当客户端退出的时候关闭连接
	defer conn.Close()

	//3、调用远程函数
	//微服务端返回的数据
	var reply AddGoodsRes
	/*
	   1、第一个参数: goods.AddGoods,goods 表示服务名称  AddGoods 方法名称
	   2、第二个参数: 给服务端的req传递数据
	   3、第三个参数: 需要传入地址,获取微服务端返回的数据
	*/
	err = conn.Call("goods.AddGoods", AddGoodsReq{
		Id:      10,
		Title:   "商品标题",
		Price:   23.5,
		Content: "商品详情",
	}, &reply)

	if err != nil {
		fmt.Println(err)
	}
	//4、获取微服务返回的数据
	fmt.Printf("%#v\n", reply)

	// 5、 调用远程GetGoods函数
	var goodsData GetGoodsRes
	err = conn.Call("goods.GetGoods", GetGoodsReq{
		Id: 12,
	}, &goodsData)
	if err != nil {
		fmt.Println(err)
	}
	//6、获取微服务返回的数据
	fmt.Printf("%#v", goodsData)
}
