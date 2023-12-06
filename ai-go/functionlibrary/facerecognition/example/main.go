/*
@File   : main.go
@Author : pan
@Time   : 2023-12-06 14:45:12
*/
package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"path/filepath"

	"github.com/Kagami/go-face"
	"gocv.io/x/gocv"
)

const (
	dataDir  = "models"
	imageDir = "images"
)

func main() {
	// 特征值
	var descriptor face.Descriptor
	// 目标距离
	matchDistance := 0.1
	// 创建一个窗口
	window := gocv.NewWindow("dlib Recognize")
	defer window.Close()

	// 颜色
	greenColor := color.RGBA{0, 255, 0, 255}
	redColor := color.RGBA{255, 0, 0, 255}

	// 加载模型
	rec, err := face.NewRecognizer(dataDir)
	if err != nil {
		log.Fatal(err)
	}
	defer rec.Close()

	// 单一脸图片
	faceImagePath := filepath.Join(imageDir, "face.jpg")
	img1 := gocv.IMRead(faceImagePath, gocv.IMReadColor)
	defer img1.Close()
	fmt.Println("正在读的单一脸图像 = ", faceImagePath)

	//
	faces, err := rec.RecognizeFile(faceImagePath)
	if err != nil {
		log.Fatalf("无法识别: %v", err)
	}
	if 0 == len(faces) {
		log.Fatal("图片没有人脸")
	}
	for _, f := range faces {
		descriptor = f.Descriptor
	}

	// 多人脸图片
	facesImagePath := filepath.Join(imageDir, "faces.jpg")
	img2 := gocv.IMRead(facesImagePath, gocv.IMReadColor)
	defer img2.Close()

	// copy bg to draw
	background := img2.Clone()
	defer background.Close()

	//
	fmt.Println("正在读的多人脸图像 = ", facesImagePath)

	//
	faces, err = rec.RecognizeFile(facesImagePath)
	if err != nil {
		log.Fatalf("无法识别: %v", err)
	}
	if 0 == len(faces) {
		log.Fatal("图片没有人脸")
	}

	for _, f := range faces {
		gocv.Rectangle(&background, f.Rectangle, redColor, 3)
		// 计算特征值之间的欧拉距离
		dist := face.SquaredEuclideanDistance(f.Descriptor, descriptor)
		fmt.Println("欧拉距离 = ", dist)
		c := redColor
		if dist < matchDistance {
			c = greenColor
		}
		// 在图片上画人脸框
		pt := image.Pt(f.Rectangle.Min.X, f.Rectangle.Min.Y-20)
		gocv.PutText(&background, "jay", pt, gocv.FontHersheyPlain, 2, c, 2)
	}

	// 显示图片
	window.IMShow(background)
	for {
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
