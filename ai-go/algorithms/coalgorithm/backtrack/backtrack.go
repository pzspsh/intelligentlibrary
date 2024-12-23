package main

// 回溯算法
import "fmt"

// 全排列
func Permute(nums []int) [][]int {
	n := len(nums)
	ans := [][]int{}
	temp := []int{}
	var backTrack func(selected []int)
	var contain = func(arr []int, x int) bool {
		for i := 0; i < len(arr); i++ {
			if x == arr[i] {
				return true
			}
		}
		return false
	}
	// 回溯函数
	backTrack = func(selected []int) {
		if len(selected) == n {
			selectedCopy := make([]int, n)
			copy(selectedCopy, selected)
			ans = append(ans, selectedCopy)
			return
		}
		for i := 0; i < n; i++ {
			if contain(selected, nums[i]) {
				continue
			}
			selected = append(selected, nums[i])
			backTrack(selected)
			selected = selected[0 : len(selected)-1]
		}
	}
	backTrack(temp)
	return ans
}

func main() {
	nums := []int{1, 2, 3, 4}
	ans := Permute(nums)
	fmt.Println("最终结果：", ans)
}

/*
最终结果： [
	[1 2 3 4]
	[1 2 4 3]
	[1 3 2 4]
	[1 3 4 2]
	[1 4 2 3]
	[1 4 3 2]
	[2 1 3 4]
	[2 1 4 3]
	[2 3 1 4]
	[2 3 4 1]
	[2 4 1 3]
	[2 4 3 1]
	[3 1 2 4]
	[3 1 4 2]
	[3 2 1 4]
	[3 2 4 1]
	[3 4 1 2]
	[3 4 2 1]
	[4 1 2 3]
	[4 1 3 2]
	[4 2 1 3]
	[4 2 3 1]
	[4 3 1 2]
	[4 3 2 1]
]
*/
