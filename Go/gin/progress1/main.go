package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//加载静态文件
	r.Static("/xxx", "./static1")
	r.Static("static", "./static")

	r.LoadHTMLGlob("template/*")
	//渲染模板
	r.GET("/deal", func(c *gin.Context) {
		c.HTML(http.StatusOK, "deal.html", nil)
	})
	r.GET("/vmled", func(c *gin.Context) {
		c.HTML(http.StatusOK, "vmled.html", nil)
	})
	//运行
	r.Run(":8080")
}
