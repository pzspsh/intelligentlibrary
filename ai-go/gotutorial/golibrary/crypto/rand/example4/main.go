/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 11:01:36
*/
package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

/*
返回一个在[0, max)区间服从均匀分布的随机值，如果max<=0则会panic。
*/
func main() {
	//从128开始，这样就能够将(max.BitLen() % 8) == 0的情况包含在里面
	for n := 128; n < 140; n++ {
		b := new(big.Int).SetInt64(int64(n)) //将new(big.Int)设为int64(n)并返回new(big.Int)
		fmt.Printf("max Int is : %v\n", b)
		i, err := rand.Int(rand.Reader, b)
		if err != nil {
			fmt.Printf("Can't generate random value: %v, %v", i, err)
		}
		fmt.Printf("rand Int is : %v\n", i)
	}
}
