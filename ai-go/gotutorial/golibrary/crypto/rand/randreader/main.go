/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 11:02:28
*/
package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	for n := 2; n < 10; n++ {
		p, err := rand.Prime(rand.Reader, n) //n代表位数，比如3为2位，127为7位
		if err != nil {
			fmt.Printf("Can't generate %d-bit prime: %v", n, err)
		}
		if p.BitLen() != n { //返回p的绝对值的字位数，0的字位数为0
			fmt.Printf("%v is not %d-bit", p, n)
		}
		if !p.ProbablyPrime(32) { //对p进行32次Miller-Rabin质数检测。如果方法返回真则p是质数的几率为1-(1/4)**32；否则p不是质数
			fmt.Printf("%v is not prime", p)
		}
		fmt.Println(p)
	}
}
