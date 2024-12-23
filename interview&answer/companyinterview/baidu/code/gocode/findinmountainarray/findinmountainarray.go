/*
@File   : findinmountainarray.go
@Author : pan
@Time   : 2023-05-24 11:18:09
*/
package main

import "fmt"

type MountainArray struct {
	arr []int
}

func (m *MountainArray) get(index int) int {
	return m.arr[index]
}
func (m *MountainArray) length() int {
	return len(m.arr)
}

func main() {
	m := MountainArray{[]int{1, 5, 2}}
	fmt.Println(findInMountainArray(1, &m))
}
func findInMountainArray(target int, mountainArr *MountainArray) int {
	topIndex := findMountainTop(mountainArr)
	LeftIndex := findInInc(0, topIndex+1, mountainArr, target)
	if LeftIndex >= 0 {
		return LeftIndex
	} else {
		RightIndex := findInDes(topIndex+1, mountainArr.length(), mountainArr, target)
		if RightIndex >= 0 {
			return RightIndex
		} else {
			return -1
		}
	}
	// return -2
}

func findInInc(l, r int, mountainArr *MountainArray, target int) int {
	if ((target - mountainArr.get(l)) * (target - mountainArr.get(r-1))) > 0 {
		return -1
	}
	for l < r {
		m := (l + r) / 2
		if mountainArr.get(m) > target {
			r = m
		} else if mountainArr.get(m) < target {
			l = m + 1
		} else {
			return m
		}
	}
	return -1
}
func findInDes(l, r int, mountainArr *MountainArray, target int) int {
	if ((target - mountainArr.get(l)) * (target - mountainArr.get(r-1))) > 0 {
		return -1
	}
	for l < r {
		m := (l + r) / 2
		if mountainArr.get(m) > target {
			l = m + 1
		} else if mountainArr.get(m) < target {
			r = m
		} else {
			return m
		}
	}
	return -1
}
func findMountainTop(mountainArr *MountainArray) int {
	l, r := 1, mountainArr.length()-1
	for l < r-1 {
		m := (l + r) / 2
		if mountainArr.get(m) > mountainArr.get(m-1) {
			l = m
		} else {
			r = m
		}
	}
	return l
}
