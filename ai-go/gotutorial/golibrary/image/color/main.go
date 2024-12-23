/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 11:04:25
*/
package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strconv"
	"strings"

	"github.com/liujiawm/graphics-go/graphics"
)

// 该工具支持将图片色彩反转，图片灰化，图片转为字符画。
// author iccboy 2017-9-2
func main() {
	args := os.Args //获取用户输入的所有参数
	if args == nil || len(args) < 4 || !(args[1] == "-r" || args[1] == "-g" || args[1] == "-t" || args[1] == "-c") {
		usage()
		return
	}
	fmt.Print("...转换中...")
	option := args[1]
	source := args[2]
	target := args[3]

	ff, _ := os.ReadFile(source)
	bbb := bytes.NewBuffer(ff)
	m, _, _ := image.Decode(bbb)
	if option == "-r" {
		newRgba := fzImage(m)
		f, _ := os.Create(target)
		defer f.Close()
		encode(source, f, newRgba)
	} else if option == "-g" {
		newGray := hdImage(m)
		f, _ := os.Create(target)
		defer f.Close()
		encode(source, f, newGray)
	} else if option == "-c" {
		rectWidth := 200
		if len(args) > 4 {
			rectWidth, _ = strconv.Atoi(args[4])
		}
		newRgba := rectImage(m, rectWidth)
		f, _ := os.Create(target)
		defer f.Close()
		encode(source, f, newRgba)
	} else {
		ascllimage(m, target)
	}
	fmt.Println("转换完成...")
}

// 帮助提示信息
var usage = func() {
	fmt.Println("输入错误，请按照下面的格式输入：")
	fmt.Println("使用: imagetool [OPTION]  source_image [output]")
	fmt.Println("  Options is flow:")
	fmt.Println("    -r         图片颜色翻转")
	fmt.Println("    -g         图片灰度")
	fmt.Println("    -c         缩放文本，该参数时，可以传入图片缩放的宽度 如：imagetool -c 1.jpg c1.jpg 100")
	fmt.Println("    -t         转成文本")

}

// 图片编码
func encode(inputName string, file *os.File, rgba *image.RGBA) {
	if strings.HasSuffix(inputName, "jpg") || strings.HasSuffix(inputName, "jpeg") {
		jpeg.Encode(file, rgba, nil)
	} else if strings.HasSuffix(inputName, "png") {
		png.Encode(file, rgba)
	} else if strings.HasSuffix(inputName, "gif") {
		gif.Encode(file, rgba, nil)
	} else {
		fmt.Println("不支持的图片格式")
	}
}

// 图片色彩反转
func fzImage(m image.Image) *image.RGBA {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	newRgba := image.NewRGBA(bounds)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			colorRgb := m.At(i, j)
			r, g, b, a := colorRgb.RGBA()
			r_uint8 := uint8(r >> 8)
			g_uint8 := uint8(g >> 8)
			b_uint8 := uint8(b >> 8)
			a_uint8 := uint8(a >> 8)
			r_uint8 = 255 - r_uint8
			g_uint8 = 255 - g_uint8
			b_uint8 = 255 - b_uint8
			newRgba.SetRGBA(i, j, color.RGBA{r_uint8, g_uint8, b_uint8, a_uint8})
		}
	}
	return newRgba
}

// 图片灰化处理
func hdImage(m image.Image) *image.RGBA {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	newRgba := image.NewRGBA(bounds)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			colorRgb := m.At(i, j)
			_, g, _, a := colorRgb.RGBA()
			g_uint8 := uint8(g >> 8)
			a_uint8 := uint8(a >> 8)
			newRgba.SetRGBA(i, j, color.RGBA{g_uint8, g_uint8, g_uint8, a_uint8})
		}
	}
	return newRgba
}

// 图片缩放, add at 2018-9-12
func rectImage(m image.Image, newdx int) *image.RGBA {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	newRgba := image.NewRGBA(image.Rect(0, 0, newdx, newdx*dy/dx))
	graphics.Scale(newRgba, m)
	return newRgba
}

// 图片转为字符画（简易版）
func ascllimage(m image.Image, target string) {
	if m.Bounds().Dx() > 300 {
		m = rectImage(m, 300)
	}
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	arr := []string{"M", "N", "H", "Q", "$", "O", "C", "?", "7", ">", "!", ":", "–", ";", "."}

	fileName := target
	dstFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			colorRgb := m.At(j, i)
			_, g, _, _ := colorRgb.RGBA()
			avg := uint8(g >> 8)
			num := avg / 18
			dstFile.WriteString(arr[num])
			if j == dx-1 {
				dstFile.WriteString("\n")
			}
		}
	}
}
