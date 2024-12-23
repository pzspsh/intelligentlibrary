# 常见算法(common algorithms)
### 六大常用算法
#### 1、枚举算法
```go

```
#### 2、动态规划
```go
案例：实现动态规划算法(背包问题)
package main

import "fmt"

const (
	// 行
	RAW int = 4
	// 列
	COL int = 5
)

// 物品的重量 舍弃数组的0位置元素
var weight = [RAW]int{0, 1, 4, 3}
// 物品的价值 舍弃数组的0位置元素
var value = [RAW]int{0, 1500, 3000, 2000}
// 动态规划网格
var cells [RAW][COL]int
// 用于回溯选择的商品 selected[i]=1 表示放进背包，0表示不放进背包
var selected [RAW]int

// 动态规划计算网格
func dynamic() {
	for i := 1; i < len(value); i++ {
		for j := 1; j < 5; j++ {
			cells[i][j] = maxValue(i, j)
		}
	}
	for i := 0; i < RAW; i++ {
		fmt.Printf("raw is %v \n", cells[i])
	}
	findBack()
	fmt.Printf("selected goods is %v \n", selected)

}

// 判断当前单元格最大价值方法
func maxValue(i, j int) int {
	// 当前商品无法放入背包，返回当前背包所能容纳的最大价值
	if j < weight[i] {
		return cells[i-1][j]
	}
	// 可放进背包时候，计算放入当前商品后的最大价值
	curr := value[i] + cells[i-1][j-weight[i]]
	// 与当前商品不放进背包的最大价值比较，看是否应该放进背包
	if curr >= cells[i-1][j] {
		return curr
	}
	return cells[i-1][j]
}

// 回溯选择的商品方法
func findBack() {
	col := COL - 1
	for i := RAW - 1; i > 0; i-- {
		if cells[i][col] > cells[i-1][col] {
			selected[i] = 1
			col = col - weight[i]
		} else {
			selected[i] = 0
		}
	}
}
```

#### 3、分治算法
```go
将一个复杂的问题分成两个或更多个相同或相似的子问题，求得子问题的解，再合并就得到原问题的解。其核心思想是“分而治之”。
##### 基本思路
- 分：将一个问题分解成多个相似且足够小的子问题。
- 治：通常通过递归的方式，对子问题进行求解。
- 合并：将子问题的解合并得到原问题的解。
```

#### 4、递归算法
```go
递归就是不断地调用函数本身。
package main
import "fmt"


func LaF(n, a1, a2 int) int {
    if n == 0 {
        return a1
    }
    return LaF(n-1, a2, a1+a2)
}

func main() {
    fmt.Println(LaF(1, 1, 1))
    fmt.Println(LaF(2, 1, 1))
    fmt.Println(LaF(3, 1, 1))
    fmt.Println(LaF(4, 1, 1))
    fmt.Println(LaF(5, 1, 1))
}
```

#### 5、回溯算法
```go
案例：针对解决的问题就是全排列
例如给个123
输出：123 132 213 231 312 321
package main

import "fmt"

func main() {
	//全排列
	input := []int{5, 4, 6, 2}
	n := len(input)
	output := [][]int{}
	var traceback func(count int, tmp []int, had map[int]bool)
	traceback = func(count int, tmp []int, had map[int]bool) {
		if count == n {
		//这里有个坑，因为是同一个tmp,所有后面的值会把前面的值给更新掉。
		//因此要创建新的变量，复制一下再apppend进去
            new:=[]int{}
            new=append(new,tmp...)
			output = append(output, new)
			return
		}
		for _, v := range input {
			if !had[v] {
				tmp = append(tmp, v)
				had[v] = true
				traceback(count+1, tmp, had)
				delete(had, v)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	traceback(0, []int{}, map[int]bool{})
	fmt.Println(output)
}
```

#### 6、贪心算法
```go
案例：最大子序和 
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
示例:
输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

package main

import (
    "fmt"
)

func main(){
    sum := MaxSubArray([]int{2,-1,3,6,7,4,-5})
    fmt.Println(sum)
}
func MaxSubArray(nums []int) int {
	if len(nums) == 1{
		return nums[0]
	}
	var currSum, maxSum = 0, nums[0]
    for _, v := range nums {
        if currSum > 0 {
            currSum += v
        } else {
            currSum = v
        }
        if maxSum < currSum {
            maxSum = currSum
        }
    }
	return maxSum
}
```
