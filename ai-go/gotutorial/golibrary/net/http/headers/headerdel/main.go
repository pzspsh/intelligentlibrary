/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:14:42
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	header := http.Header{"Date": {"1994-11-06T08:49:37Z00:00"}}
	fmt.Println(header.Get("Date"))         //1994-11-06T08:49:37Z00:00
	fmt.Println(header.Get("Content-Type")) //因为没有该字段，返回为空

	header.Set("Content-Type", "text/plain; charset=UTF-8")          //设置"Content-Type"字段
	fmt.Println(header.Get("Content-Type"))                          //返回text/plain; charset=UTF-8
	header.Set("Content-Type", "application/x-www-form-urlencoded;") //覆盖原先的值，返回application/x-www-form-urlencoded;
	fmt.Println(header.Get("Content-Type"))

	header.Add("Content-Type", "charset=UTF-8") //在"Content-Type"字段中追加值
	fmt.Println(header)                         //map[Date:[1994-11-06T08:49:37Z00:00] Content-Type:[application/x-www-form-urlencoded; charset=UTF-8]]，可见添加进去
	fmt.Println(header.Get("Content-Type"))     //但是这样获取是返回值仍是application/x-www-form-urlencoded;

	header.Del("Content-Type")              //删除该字段
	fmt.Println(header.Get("Content-Type")) //然后返回又为空
}
