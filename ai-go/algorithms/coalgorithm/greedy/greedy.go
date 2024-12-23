package main

import "sort"

// 贪心算法
func main() {

}

func Greedy(g []int, s []int) int {
	sort.Ints(g) // 内部用快速排序实现,升序
	sort.Ints(s)

	n1 := len(g)
	n2 := len(s)
	if n1 == 0 || n2 == 0 {
		return 0
	}
	n1--
	n2--
	count := 0
	for n1 >= 0 && n2 >= 0 {
		if s[n2] >= g[n1] {
			count++
			n1--
			n2--
		} else { // s[n2] < g[n1]
			n1--
		}
	}
	return count
}
