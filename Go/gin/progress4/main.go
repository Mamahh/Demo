package main

import (
	"fmt"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 默认大小限制 8M
	// r.MaxMultipartMemory = 8 << 20
	r.LoadHTMLGlob("./static/*")

	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("upload", func(c *gin.Context) {
		//接收文件
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			//保存文件
			dst := path.Join("./tmpfile", file.Filename)
			c.SaveUploadedFile(file, dst)
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("'%s' upload OK !", file.Filename),
			})
		}
	})

	r.Run(":9090")
}
