package main

import "fmt"

func main() {
	fmt.Println(numTrees1(3))
}

/*
C0 = 1
Cn+1 = 2(2n+1)/(n+2) * Cn
*/
func numTrees1(n int) int {
	c := 1
	for i := 1; i < n; i++ {
		c = c * 2 * (2*i + 1) / (i + 2)
	}
	return c
}
