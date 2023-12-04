/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 15:42:19
*/
package main

import (
	gt "github.com/mangenotwork/gathertool"
)

func main() {
	gt.SocketProxy(":8111")
}
