/*
@File   : main.go
@Author : pan
@Time   : 2024-01-10 14:08:14
*/
package main

import (
	"fmt"
	_ "image"
	"image/color"
	"log"

	"gocv.io/x/gocv"
)

/*
请确保在运行代码之前将haarcascade_frontalface_default.xml文件下载到您的工作目录中。
该文件包含了人脸识别的级联分类器。
*/

func main() {
	webcam, err := gocv.VideoCaptureDevice(0) // 打开摄像头
	if err != nil {
		log.Fatalf("Error opening webcam: %v", err)
	}
	defer webcam.Close()

	classifier := gocv.NewCascadeClassifier() // 加载人脸分类器
	defer classifier.Close()

	if !classifier.Load("haarcascade_frontalface_default.xml") {
		log.Fatalf("Error reading cascade file: haarcascade_frontalface_default.xml")
	}

	window := gocv.NewWindow("Face Detect") // 打开窗口以显示视频
	defer window.Close()

	img := gocv.NewMat() // 创建一个图像矩阵以保存帧
	defer img.Close()

	fmt.Printf("Press ESC to stop\n")

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed\n")
			return
		}
		if img.Empty() {
			continue
		}
		gray := gocv.NewMat() // 转换图像为灰度
		defer gray.Close()
		gocv.CvtColor(img, &gray, gocv.ColorBGRToGray)

		rects := classifier.DetectMultiScale(gray) // 探测人脸
		for _, r := range rects {
			gocv.Rectangle(&img, r, color.RGBA{0, 255, 0, 0}, 3) // 在原图上画矩形
		}
		window.IMShow(img) // 显示图像
		if window.WaitKey(1) == 27 {
			break
		}
	}
}
