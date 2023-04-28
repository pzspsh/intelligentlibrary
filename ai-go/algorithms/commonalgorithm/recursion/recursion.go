package main

import "fmt"

// 递归算法
func Recursion(n int) int {
	if n == 0 {
		return 1
	}
	return n * Recursion(n-1)
}

func main() {
	fmt.Println(Recursion(5))
}

/*
Recursion(5)
{5 * Recursion(4)}
{5 * {4 * Recursion(3)}}
{5 * {4 * {3 * Recursion(2)}}}
{5 * {4 * {3 * {2 * Recursion(1)}}}}
{5 * {4 * {3 * {2 * 1}}}}
{5 * {4 * {3 * 2}}}
{5 * {4 * 6}}
{5 * 24} == 120
*/
