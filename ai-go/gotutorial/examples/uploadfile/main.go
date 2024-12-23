/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 14:01:56
*/
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func uploadFile(c *gin.Context) {
	//form表单
	//c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, int64(2<<20))
	file, fileHeader, err := c.Request.FormFile("upload")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("上传文件失败: %s", err.Error()))
		return
	}

	filename := fileHeader.Filename
	out, err := os.Create(filename)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("创建文件: %s", err.Error()))
		return
	}

	defer out.Close()
	count := 0
	for {
		buf := make([]byte, 10000)
		n, err := file.Read(buf)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("读取失败: %s", err.Error()))
			return
		}
		if n == 0 {
			break
		}
		count = count + n
		out.Write(buf)
		fmt.Println(count, float64(fileHeader.Size))
		progress := float64(count) / float64(fileHeader.Size) * 100
		fmt.Printf(fmt.Sprintf("%.2f%%", progress))
	}

	c.String(http.StatusCreated, "上传成功 \n")
}

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 2 * 1024 //2M Byte，默认32M

	//路由:http://localhost:8080/upload
	router.POST("/upload", uploadFile)
	fmt.Println(router.MaxMultipartMemory)
	router.Run(":8080")
}
