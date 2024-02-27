package views

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// 定义LoginForm结构体，并使用自定义验证
type LoginForm struct {
	Username string `form:"username" binding:"required,usernamePrefix"`
	Password string `form:"password" binding:"required,min=8"`
}

// 自定义验证器的函数签名
func usernamePrefixValidator(fl validator.FieldLevel) bool {
	return strings.HasPrefix(fl.Field().String(), "prefix_")
}

func Main() {
	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册自定义验证器
		v.RegisterValidation("usernamePrefix", usernamePrefixValidator)
	}
	router.POST("/login", func(c *gin.Context) {
		var form LoginForm
		// 绑定和验证
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 处理登录逻辑...
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})
	router.Run(":8080")
}
