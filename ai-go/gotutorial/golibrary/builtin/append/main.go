package main

import (
	"fmt"
)

func main() {
	//其实考虑一下，这也是很正常的，因为在go中[]byte和string是可以直接相互转换的．
	slice := append([]byte("hello "), "world"...)
	fmt.Println(string(slice))
}
