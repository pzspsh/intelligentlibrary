/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 11:45:34
*/
package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/liujiawm/graphics-go/graphics"
)

func main() {
	src, err := LoadImage("filepath/348.png")
	if err != nil {
		log.Fatal(err)
	}

	// 缩略图的大小
	dst := image.NewRGBA(image.Rect(0, 0, 20, 80))

	// 产生缩略图,等比例缩放
	err = graphics.Scale(dst, src)
	if err != nil {
		log.Fatal(err)
	}

	// 需要保存的文件
	imgcounter := 734
	saveImage(fmt.Sprintf("%03d.png", imgcounter), dst)
}

// LoadImage decodes an image from a file.
func LoadImage(path string) (img image.Image, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	img, _, err = image.Decode(file)
	return
}

// 保存Png图片
func saveImage(path string, img image.Image) (err error) {
	// 需要保存的文件
	imgfile, err := os.Create(path)
	if err != nil {
		fmt.Println("os Create error:", err)
	}
	defer imgfile.Close()

	// 以PNG格式保存文件
	err = png.Encode(imgfile, img)
	if err != nil {
		log.Fatal(err)
	}
	return
}
