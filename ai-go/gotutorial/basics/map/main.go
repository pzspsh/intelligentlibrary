package main

import (
	"fmt"
	"maps"
)

func MapDemo1() {
	// 你可以通过遍历其中一个map的键值对，然后将它们添加到新的map中，最后再添加第二个map的键值对来实现合并。
	map1 := map[string]int{
		"a": 1,
		"b": 2,
	}
	map2 := map[string]int{
		"c": 3,
		"d": 4,
	}

	// 创建一个新的map来合并map1和map2
	mergedMap := make(map[string]int)
	for key, value := range map1 {
		mergedMap[key] = value
	}
	for key, value := range map2 {
		mergedMap[key] = value
	}

	fmt.Println(mergedMap) // 输出: map[a:1 b:2 c:3 d:4]
}

func MapDemo2() {
	// 更新一个已有的map（如果键冲突，第二个map的值将覆盖第一个map的值）
	map1 := map[string]int{
		"a": 1,
		"b": 2,
	}
	map2 := map[string]int{
		"b": 3, // 这里会覆盖map1中的"b": 2
		"c": 4,
	}

	// 使用range遍历第二个map，更新第一个map的值
	for key, value := range map2 {
		map1[key] = value // 如果key在map1中不存在，则会被添加；如果存在，则会被更新为map2中的值
	}

	fmt.Println(map1) // 输出: map[a:1 b:3 c:4]
}

func MapDemo3() {
	/*
		注意：使用maps.Copy会直接修改原始的map1。如果你想保留原始的map1不变，可以先使用maps.Clone复制一份，
		然后再用maps.Copy合并到这个副本上。例如：mergedMap := maps.Clone(map1); maps.Copy(mergedMap, map2)。
		这样，mergedMap就是合并后的结果，而map1保持不变。
	*/
	map1 := map[string]int{
		"a": 1,
		"b": 2,
	}
	map2 := map[string]int{
		"c": 3,
		"d": 4,
		"b": 5,
	}
	maps.Copy(map2, map1) // 使用maps.Copy来合并两个map，注意这将修改原始的map1，如果不需要修改原始map，可以先复制一份再操作。例如：mergedMap := maps.Clone(map1); maps.Copy(mergedMap, map2)
	fmt.Println(map2)     // 输出: map[a:1 b:2 c:3 d:4] 注意：这也会修改原始的map1，如果要保留原图，可以先复制一份。例如：mergedMap := maps.Clone(map1); maps.Copy(mergedMap, map2) 然后使用mergedMap。

	// mergedMap := maps.Clone(map1)
	// maps.Copy(mergedMap, map2)
	// fmt.Println(mergedMap) // 输出: map[a:1 b:5 c:3 d:4]
}

func main() {
	MapDemo3()
}
