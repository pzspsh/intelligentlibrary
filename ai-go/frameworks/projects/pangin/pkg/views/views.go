/*
@File   : views.go
@Author : pan
@Time   : 2023-06-11 22:24:57
*/
package views

import (
	"net/http"
	"pangin/pkg/models"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreatePandDemo(c *gin.Context) {
	// 前端页面填写待办事项 点击提交 会发请求到这里
	// 1. 从请求中把数据拿出来
	var pandemo models.PanDemo
	c.BindJSON(&pandemo)
	// 2. 存入数据库
	err := models.CreatePanDemo(&pandemo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, pandemo)
		//c.JSON(http.StatusOK, gin.H{
		//	"code": 2000,
		//	"msg": "success",
		//	"data": todo,
		//})
	}
}

func GetPandDemo(c *gin.Context) {
	// 查询todo这个表里的所有数据
	pandemo, err := models.GetAllPanDemo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, pandemo)
	}
}

func UpdatePandDemo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"message": "无效id"})
		return
	}
	pandemo, err := models.GetAPanDemo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.BindJSON(&pandemo)
	if err = models.UpdatePanDemo(pandemo); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, &pandemo)
	}
}

func DeletePandDemo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"message": "无效id"})
		return
	}
	if err := models.DeletePanDemo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
