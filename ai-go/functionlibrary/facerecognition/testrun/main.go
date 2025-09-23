//go:build !windows

package main

import (
	"fmt"

	"github.com/Kagami/go-face"
)

const dataDir = "testdata"

// data 目录下两个对应的文件夹目录
const (
	modelDir  = dataDir + "/models"
	imagesDir = dataDir + "/images"
)

func main() {
	fmt.Println("Face Recognition...")
	rec, err := face.NewRecognizer(modelDir) // 初始化识别器
	if err != nil {
		fmt.Println("Cannot INItialize recognizer")
	}
	defer rec.Close()
	fmt.Println("Recognizer Initialized")
}
