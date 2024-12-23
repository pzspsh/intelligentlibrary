/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 17:41:04
*/
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// 比较结构体
	p1 := Person{Name: "Alice", Age: 18}
	p2 := Person{Name: "Alice", Age: 18}
	p3 := Person{Name: "Bob", Age: 20}

	fmt.Println(p1 == p2)                  // true
	fmt.Println(p1 == p3)                  // false
	fmt.Println(reflect.DeepEqual(p1, p2)) // true
	fmt.Println(reflect.DeepEqual(p1, p3)) // false

	p1JSON, _ := json.Marshal(p1)
	p2JSON, _ := json.Marshal(p2)
	p3JSON, _ := json.Marshal(p3)
	fmt.Println(string(p1JSON) == string(p2JSON)) // true
	fmt.Println(string(p1JSON) == string(p3JSON)) // false

	// 比较切片
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{4, 5, 6}

	// fmt.Println(s1 == s2)                  // true
	// fmt.Println(s1 == s3)                  // false
	fmt.Println(reflect.DeepEqual(s1, s2)) // true
	fmt.Println(reflect.DeepEqual(s1, s3)) // false

	s1JSON, _ := json.Marshal(s1)
	s2JSON, _ := json.Marshal(s2)
	s3JSON, _ := json.Marshal(s3)
	fmt.Println(string(s1JSON) == string(s2JSON)) // true
	fmt.Println(string(s1JSON) == string(s3JSON)) // false

	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"a": 1, "b": 2, "c": 3}
	m3 := map[string]int{"a": 1, "b": 2, "c": 4}

	// fmt.Println(m1 == m2)                  // true
	// fmt.Println(m1 == m3)                  // false
	fmt.Println(reflect.DeepEqual(m1, m2)) // true
	fmt.Println(reflect.DeepEqual(m1, m3)) // false

	m1JSON, _ := json.Marshal(m1)
	m2JSON, _ := json.Marshal(m2)
	m3JSON, _ := json.Marshal(m3)
	fmt.Println(string(m1JSON) == string(m2JSON)) // true
	fmt.Println(string(m1JSON) == string(m3JSON)) // false

}
