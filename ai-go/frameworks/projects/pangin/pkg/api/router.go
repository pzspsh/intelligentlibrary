/*
@File   : router.go
@Author : pan
@Time   : 2023-06-07 14:27:43
*/
package api

import (
	"pangin/configs"
	"pangin/pkg/views"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	if configs.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "../static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("../templates/*")
	r.GET("/", views.IndexHandler)

	v1Group := r.Group("v1")
	{
		v1Group.POST("/pandemo", views.CreatePandDemo)
		v1Group.GET("/pandemo", views.GetPandDemo)
		v1Group.PUT("/pandemo/:id", views.UpdatePandDemo)
		v1Group.DELETE("pandemo/:id", views.DeletePandDemo)
	}
	return r
}
