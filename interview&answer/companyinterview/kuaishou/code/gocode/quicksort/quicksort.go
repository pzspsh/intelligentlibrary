/*
@File   : quicksort.go
@Author : pan
@Time   : 2023-05-17 10:56:08
*/
package gocode

// 快速排序
// 版本1
func QuickSortV1(arr []int, low, hight int) {
	// 当 low = hight  时跳出
	if low >= hight {
		return
	}
	left, right := low, hight
	pivot := arr[left] // 为了简单起见，直接取左边的第一个数

	for left < right {
		// 先从右边开始迭代
		// 右边的数如果比pivot大，那么就应该将他放在右边，继续向左滑动，遇到一个比他小的为止
		for left < right && arr[right] >= pivot {
			right--
		}

		// 小数移动到左边
		if left < right {
			arr[left] = arr[right]
		}

		// 左边的数如果比pivot小，那么就应该将他放在左边，继续向右滑动，遇到一个比他大的为止
		for left < right && arr[left] < pivot {
			left++
		}

		// 大数移动到又边
		if left < right {
			arr[right] = arr[left]
		}

		// left == right ,pivot 即是最终位置
		if left <= right {
			arr[left] = pivot
		}
	}

	//因为 pivot 的最终位置已锁定

	// 继续排序左边部分
	QuickSortV1(arr, low, right-1)
	// 继续排序右边部分
	QuickSortV1(arr, right+1, hight)
}

// 版本2
func QuickSortV2(arr []int, low, hight int) {
	if low < hight {
		left, right := low, hight
		pivot := arr[(low+hight)/2] // 这里的经验值取的是中间数，经过 Benchmark 测试，确实比较优秀

		for left <= right {
			// 从左边开始迭代
			// 左边的数如果比 pivot 小，那么就应该将他放在左边，继续向右滑动，遇到一个比他大的为止
			for arr[left] < pivot {
				left++
			}

			// 右边的数如果比 pivot 大，那么就应该将他放在右边，继续向左滑动，遇到一个比他小的为止
			for arr[right] > pivot {
				right--
			}

			// 这里进行一次交换，将上面碰到的大数和小数交换一次
			//left 继续右走，right 继续左走 注意这里还不一定相遇，去继续执行上面的逻辑
			if left <= right {
				arr[left], arr[right] = arr[right], arr[left]
				left++
				right--
			}
		}

		// 【 xxx[xxxxx]xxxxxx】
		// [ xxx【xxxxx】xxxxxx]
		// [] => left,right
		// 【】 => low,hight
		// 这里其实挺费解？ 但是如果 right 在 low 左侧，那么我们排序的访问就不是 low到 hight 这一段切片
		// 丧失了排序一段 quickSort 中参数的意义
		// 如果想不通也可以去掉，让递归到下一层调用栈中弹出
		if low < right {
			QuickSortV2(arr, low, right)
		}
		// 同理
		if hight > left {
			QuickSortV2(arr, left, hight)
		}
	}
}

// 版本3
func QuickSortV3(arr []int, left, right int) {
	if left >= right {
		return
	}
	cur, lo := left+1, left
	for cur <= right {
		if arr[cur] <= arr[left] {
			arr[lo+1], arr[cur] = arr[cur], arr[lo+1]
			lo++
		}
		cur++
	}
	arr[left], arr[lo] = arr[lo], arr[left]
	QuickSortV3(arr, left, lo-1)
	QuickSortV3(arr, lo+1, right)
}

// 版本4
func QuickSort3(arr []int, left, right int) {
	if left >= right {
		return
	}
	pivot := arr[left]
	lo, gt, cur := left, right+1, left+1

	for cur < gt {
		if arr[cur] < pivot {
			arr[cur], arr[lo+1] = arr[lo+1], arr[cur]
			lo++
			cur++
		} else if arr[cur] > pivot {
			arr[cur], arr[gt-1] = arr[gt-1], arr[cur]
			gt--
		} else {
			cur++
		}
	}

	arr[left], arr[lo] = arr[lo], arr[left]
	QuickSort3(arr, left, lo-1)
	QuickSort3(arr, gt, right)
}
