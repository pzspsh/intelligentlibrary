package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(minPathSum1(arr))
}

// 最小路径和
func minPathSum1(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j != 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if i != 0 && j == 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else if i != 0 && j != 0 {
				grid[i][j] = min1(grid[i-1][j], grid[i][j-1]) + grid[i][j]
			}
		}
	}
	return grid[n-1][m-1]
}

func min1(a, b int) int {
	if a > b {
		return b
	}
	return a
}
