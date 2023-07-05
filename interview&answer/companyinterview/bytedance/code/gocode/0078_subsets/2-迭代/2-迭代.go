package main

import "fmt"

func main() {
	fmt.Println(subsets1([]int{1, 2, 3}))
}

func subsets1(nums []int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int{})
	for i := 0; i < len(nums); i++ {
		temp := make([][]int, len(res))
		for key, value := range res {
			value = append(value, nums[i])
			temp[key] = append(temp[key], value...)
		}
		for _, v := range temp {
			res = append(res, v)
		}
	}
	return res
}
