package main

import "fmt"

// 动态规划算法
func main() {
	n, m, k := 0, 0, 0
	fmt.Scan(&n, &m, &k)
	data := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i][0], &data[i][1])
	}

	//DP[i][j]表示对于前i个宝箱有j次机会的最大金币数
	//边界条件DP[0][0]=data[0][1]
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k+1)

	}
	ans := 0
	dp[0][0] = data[0][1]
	for i := 1; i < n; i++ {
		for j := 1; j < k+1; j++ {
			for l := i - 1; l >= 0; l-- {
				ans = max(ans, dp[i][j])
				if data[i][0]-data[l][0] > m {
					break
				} else {
					if dp[l][j-1] == 0 {
					} else {
						dp[i][j] = max(dp[i][j], dp[l][j-1]+data[i][1])
					}
				}
			}

		}
	}
	fmt.Println(ans)
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
