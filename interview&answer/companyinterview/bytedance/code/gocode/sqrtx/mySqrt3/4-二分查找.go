package main

import "fmt"

func main() {
	fmt.Println(mySqrt3(10))
	fmt.Println(mySqrt3(1))
	fmt.Println(mySqrt3(8))
	fmt.Println(mySqrt3(0))
}

func mySqrt3(x int) int {
	left := 1
	right := x
	for left <= right {
		mid := (left + right) / 2
		if mid == x/mid {
			return mid
		} else if mid < x/mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left*left <= x {
		return left
	} else {
		return left - 1
	}
}
