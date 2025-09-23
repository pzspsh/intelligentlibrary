package main

import "fmt"

func TowArraySum(arr1, arr2 []int) []int {
	return append(arr1, arr2...)
}

func main() {
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{5, 6, 7}
	// arr2 = append(arr2, arr1...)
	result := TowArraySum(arr1, arr2)
	fmt.Println(result)
}
