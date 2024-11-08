/*
@File   : main.go
@Author : pan
@Time   : 2024-11-07 10:25:19
*/
package main

import (
	"fmt"
)

func Demo() {
	map1 := map[string]int{"key1": 10, "key2": 20}
	map2 := map[string]int{"key1": 100, "key3": 300}
	keyToCheck := "key4"
	// 判断keyToCheck是否同时存在于map1和map2
	existsInBoth := map1[keyToCheck] > 0 && map2[keyToCheck] > 0
	fmt.Println("Exists in both maps:", existsInBoth)
}

func checkKeyInMap(m map[string]int, key string) bool {
	_, exists := m[key]
	return exists
}

func Demo2() {
	map1 := map[string]int{"key1": 1, "key3": 3}
	map2 := map[string]int{"key2": 2, "key4": 4}

	key1InMap1 := checkKeyInMap(map1, "key1")
	key2InMap2 := checkKeyInMap(map2, "key2")

	fmt.Printf("key1 is in map1: %t\n", key1InMap1)
	fmt.Printf("key2 is in map2: %t\n", key2InMap2)
}

func CheckKeyInMap[T any](m map[string]T, key1, key2 string) (T, bool) {
	if value, exists := m[key1]; exists {
		return value, exists
	} else if value, exists := m[key2]; exists {
		return value, exists
	} else {
		return value, false
	}
}

type cmap struct {
	field1 string
	filed2 string
}

func main() {
	data := map[string]cmap{
		"ab": {field1: "hello", filed2: "world"},
		"cd": {field1: "hello", filed2: "world"},
	}
	v, ok := CheckKeyInMap(data, "ae", "ad")
	fmt.Println(v, ok)
}
