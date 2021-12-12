package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//解析模板
	r.LoadHTMLGlob("./template/*")

	//渲染模板
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	//处理返回数据
	r.POST("/login", func(c *gin.Context) {
		//直接获取form表单提交的数据
		username := c.PostForm("user")
		password := c.PostForm("pwd")
		//添加默认值
		// username := c.DefaultPostForm("username", "somebody")
		// password := c.DefaultPostForm("pwd", "xxx")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})

	r.Run(":9090")
}
