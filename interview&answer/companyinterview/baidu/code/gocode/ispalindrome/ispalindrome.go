/*
@File   : ispalindrome.go
@Author : pan
@Time   : 2023-05-24 10:23:08
*/
package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(12321))
}

// 回文数
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		temp := x % 10
		revertedNumber = revertedNumber*10 + temp
		x = x / 10
	}
	// for example:
	// x = 1221  => x = 12 revertedNumber = 12
	// x = 12321 => x = 12 revertedNumber = 123
	return x == revertedNumber || x == revertedNumber/10
}
