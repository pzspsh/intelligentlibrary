/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 11:01:07
*/
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"
)

func main() {
	const width = 130
	const height = 50

	im := image.NewGray(image.Rectangle{Max: image.Point{X: width, Y: height}})
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			dist := math.Sqrt(math.Pow(float64(x-width/2), 2)/3+math.Pow(float64(y-height/2), 2)) / (height / 1.5) * 255
			var gray uint8
			if dist > 255 {
				gray = 255
			} else {
				gray = uint8(dist)
			}
			im.SetGray(x, y, color.Gray{Y: 255 - gray})
		}
	}
	pi := image.NewPaletted(im.Bounds(), []color.Color{
		color.Gray{Y: 255},
		color.Gray{Y: 160},
		color.Gray{Y: 70},
		color.Gray{Y: 35},
		color.Gray{Y: 0},
	})

	draw.FloydSteinberg.Draw(pi, im.Bounds(), im, image.ZP)
	shade := []string{" ", "░", "▒", "▓", "█"}
	for i, p := range pi.Pix {
		fmt.Print(shade[p])
		if (i+1)%width == 0 {
			fmt.Print("\n")
		}
	}
}
