/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 12:02:25
*/
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

const (
	dx = 256
	dy = 256
)

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dx)
	for i := range pic {
		pic[i] = make([]uint8, dy)
		for j := range pic[i] {
			pic[i][j] = uint8(i * j % 255)
		}
	}
	return pic
}
func main() {
	file, err := os.Create("filepath/test.png")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	rgba := image.NewRGBA(image.Rect(0, 0, dx, dy))
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			rgba.Set(x, y, color.RGBA{uint8(x * y % 255), uint8(x * y % 255), 0, 255})
		}
	}
	err = png.Encode(file, rgba)
	if err != nil {
		fmt.Println(err)
	}
}
