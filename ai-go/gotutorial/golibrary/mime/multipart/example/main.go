/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 15:23:21
*/
package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// 需要上传的文件路径
	filePath := "path/to/file.jpg"

	// 打开要上传的文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}
	defer file.Close()

	// 创建multipart.Writer，用于构造multipart/form-data格式的请求体
	var requestBody bytes.Buffer
	multipartWriter := multipart.NewWriter(&requestBody)

	// 创建一个multipart.Part，用于表示文件字段
	part, err := multipartWriter.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		fmt.Println("Failed to create form file:", err)
		return
	}

	// 将文件内容复制到multipart.Part中
	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println("Failed to copy file content:", err)
		return
	}

	// 添加其他表单字段
	multipartWriter.WriteField("title", "My file")

	// 关闭multipart.Writer，以便写入Content-Type和boundary
	err = multipartWriter.Close()
	if err != nil {
		fmt.Println("Failed to close multipart writer:", err)
		return
	}

	// 创建HTTP请求
	req, err := http.NewRequest("POST", "http://example.com/upload", &requestBody)
	if err != nil {
		fmt.Println("Failed to create request:", err)
		return
	}

	// 设置Content-Type为multipart/form-data
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	// 发送HTTP请求
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return
	}
	defer resp.Body.Close()

	// 处理响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response:", err)
		return
	}
	fmt.Println("Response:", string(respBody))
}
