/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 14:43:12
*/
package main

import (
	"fmt"
	"math"
)

// 定义了interface：几何体，以及方法：计算面积、计算周长
// 定义顺序：type、interface名、interface
type geometry interface {
	area() float64
	perim() float64
}

// 定义了struct：矩形和圆形，分别实现了geometry定义的所有接口
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// 矩形和圆形都可以作为几何体类型的参数传递
	measure(r)
	measure(c)
}
