/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 15:36:41
*/
package main

import (
	"fmt"
	"testing"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 测试用例需要以Test开头
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {

		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		// 用Run启动子测试用例，每个用例用数据表中获取输入
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// 基准测试需要以Benchmark开头
func BenchmarkIntMin(b *testing.B) {
	// b.N是执行次数，由底层自动决定，以产生稳定精确的性能统计
	for i := 0; i < b.N; i++ {
		re := IntMin(1, 2)
		fmt.Println(re)
	}
}
