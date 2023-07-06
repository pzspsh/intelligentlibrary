package main

import "fmt"

func main() {
	fmt.Println(numTrees2(3))
}

func numTrees2(n int) int {
	c := 1
	for i := 1; i <= n; i++ {
		c = c * (n + i) / i
	}
	return c / (n + 1)
}
