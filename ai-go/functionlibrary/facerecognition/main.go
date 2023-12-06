/*
@File   : main.go
@Author : pan
@Time   : 2023-06-06 14:58:23
*/
package main

import (
	"fmt"
	"log"

	"github.com/Kagami/go-face"
)

const dataDir = "testdata"

// testdata 目录下两个对应的文件夹目录
const (
	modelDir  = dataDir + "/models"
	imagesDir = dataDir + "/images"
)

// 图片中的人名
var labels = []string{
	"萧敬腾",
	"周杰伦",
	"unknow",
	"王力宏",
	"陶喆",
	"林俊杰",
}

func main() {
	fmt.Println("Face Recognition...")

	// 初始化识别器
	rec, err := face.NewRecognizer(modelDir)
	if err != nil {
		fmt.Println("Cannot INItialize recognizer")
	}
	defer rec.Close()

	fmt.Println("Recognizer Initialized")

	// 调用该方法，传入路径。返回面部数量和任何错误
	faces, err := rec.RecognizeFile("heyin.jpeg")
	if err != nil {
		log.Fatalf("无法识别: %v", err)
	}
	// 打印人脸数量
	fmt.Println("图片人脸数量: ", len(faces))

	var samples []face.Descriptor
	var peoples []int32
	for i, f := range faces {
		samples = append(samples, f.Descriptor)
		// 每张脸唯一 id
		peoples = append(peoples, int32(i))
	}

	// 传入样例到识别器
	rec.SetSamples(samples, peoples)

	RecognizePeople(rec, "jay.jpeg")
	RecognizePeople(rec, "linjunjie.jpeg")
	RecognizePeople(rec, "taozhe.jpeg")
}

func RecognizePeople(rec *face.Recognizer, file string) {
	people, err := rec.RecognizeSingleFile(file)
	if err != nil {
		log.Fatalf("无法识别: %v", err)
	}
	if people == nil {
		log.Fatalf("图片上不是一张脸")
	}
	peopleID := rec.Classify(people.Descriptor)
	if peopleID < 0 {
		log.Fatalf("无法区分")
	}
	fmt.Println(peopleID)
	fmt.Println(labels[peopleID])
}
