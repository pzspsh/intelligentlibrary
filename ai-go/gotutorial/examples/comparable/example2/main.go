/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 11:13:05
*/
package main

import (
	"fmt"
)

// Set是元素的集合
type Set[T comparable] map[T]struct{}

// NewSet返回一组具有指定类型的元素
func NewSet[T comparable](es ...T) Set[T] {
	s := Set[T]{}
	for _, e := range es {
		s.Add(e)
	}
	return s
}

// Len报告s的元素个数
func (s *Set[T]) Len() int {
	return len(*s)
}

// IsEmpty报告是否为空
func (s *Set[T]) IsEmpty() bool {
	return s.Len() == 0
}

// Add 添加元素到集合s
// if element is already in s this has no effect
// 如果元素已经在s中了，这没有影响
func (s *Set[T]) Add(es ...T) {
	for _, e := range es {
		(*s)[e] = struct{}{}
	}
}

// Remove 从集合中删除元素
// if element is not in s this has no effect
func (s *Set[T]) Remove(es ...T) {
	for _, e := range es {
		delete(*s, e)
	}
}

// Contains 报告v是否在s中
func (s *Set[T]) Contains(v T) bool {
	_, ok := (*s)[v]
	return ok
}

// Clone 创建一个具有与s相同元素的新集合
func (s *Set[T]) Clone() Set[T] {
	r := Set[T]{}
	r.Add(s.ToSlice()...)
	return r
}

// ToSlice 将集合变换为切片
func (s *Set[T]) ToSlice() []T {
	r := make([]T, 0, s.Len())

	for e := range *s {
		r = append(r, e)
	}

	return r
}

// Union 某些集合的并集(如:A、B)
// is the set of elements that are in either A or B
func Union[T comparable](sets ...Set[T]) Set[T] {
	r := NewSet[T]()
	for _, s := range sets {
		r.Add(s.ToSlice()...)
	}
	return r
}

func main() {
	s := NewSet("a", "b", "c")
	s.Add("c", "d", "e")
	s.Remove("b", "d")
	fmt.Printf("set: %v \n", s.ToSlice())

	s = Union(s, NewSet("e", "f"), NewSet("a", "f"))
	fmt.Printf("set: %v \n", s.ToSlice())
}
