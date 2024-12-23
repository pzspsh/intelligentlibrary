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
	var descriptor face.Descriptor             // 特征值
	matchDistance := 0.1                       // 目标距离
	window := gocv.NewWindow("dlib Recognize") // 创建一个窗口
	defer window.Close()
	greenColor := color.RGBA{0, 255, 0, 255} // 颜色
	redColor := color.RGBA{255, 0, 0, 255}
	rec, err := face.NewRecognizer(dataDir) // 加载模型
	if err != nil {
		log.Fatal(err)
	}
	defer rec.Close()

	faceImagePath := filepath.Join(imageDir, "face.jpg") // 单一脸图片
	img1 := gocv.IMRead(faceImagePath, gocv.IMReadColor)
	defer img1.Close()
	fmt.Println("正在读的单一脸图像 = ", faceImagePath)

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

	facesImagePath := filepath.Join(imageDir, "faces.jpg") // 多人脸图片
	img2 := gocv.IMRead(facesImagePath, gocv.IMReadColor)
	defer img2.Close()

	background := img2.Clone() // copy bg to draw
	defer background.Close()
	fmt.Println("正在读的多人脸图像 = ", facesImagePath)

	faces, err = rec.RecognizeFile(facesImagePath)
	if err != nil {
		log.Fatalf("无法识别: %v", err)
	}
	if 0 == len(faces) {
		log.Fatal("图片没有人脸")
	}

	for _, f := range faces {
		gocv.Rectangle(&background, f.Rectangle, redColor, 3)

		dist := face.SquaredEuclideanDistance(f.Descriptor, descriptor) // 计算特征值之间的欧拉距离
		fmt.Println("欧拉距离 = ", dist)
		c := redColor
		if dist < matchDistance {
			c = greenColor
		}

		pt := image.Pt(f.Rectangle.Min.X, f.Rectangle.Min.Y-20) // 在图片上画人脸框
		gocv.PutText(&background, "jay", pt, gocv.FontHersheyPlain, 2, c, 2)
	}

	window.IMShow(background) // 显示图片
	for {
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
