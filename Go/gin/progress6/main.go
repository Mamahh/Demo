package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func indexhandle(c *gin.Context) {
	//取值
	name, ok := c.Get("name")
	if !ok {
		name = "匿名用户"
	}
	c.JSON(http.StatusOK, gin.H{
		"message": name,
	})
}

//中间件函数
func m1(c *gin.Context) {
	fmt.Println("start m1...")
	//计时
	start := time.Now()
	// go funcXX(c.Copy()) //在funcxx()中只能使用c的拷贝,以防c被修改导致运行异常
	c.Next() //调用后续的处理函数
	// c.Abort() //中止后续的函数
	cost := time.Since(start)
	fmt.Printf("cost time is %v\n", cost)
	fmt.Println("end m1...")
}

//中间件函数
func m2(c *gin.Context) {
	fmt.Println("start m2...")
	c.Set("name", "mamahh") //传值
	c.Next()                //调用后续的处理函数
	// c.Abort() 		//中止后续的处理函数
	fmt.Println("end m2...")
}

//登陆中间件
func auth(doCheck bool) gin.HandlerFunc {
	//闭包
	//连接数据库
	//或者其他准备工作
	return func(c *gin.Context) {
		if doCheck {
			//存放具体的逻辑
			//是否登陆
			//if 登陆
			//c.Next()
			//else
			//c.Abort()
		} else {
			c.Next()
		}
	}
}

func main() {
	r := gin.Default() // 默认使用Logger(),Recovery() 中间件

	// 全局注册中间件函数
	r.Use(m1, m2, auth(true))

	// 单独注册中间件函数
	// r.GET("/index", m1, indexhandle)

	// 路由组注册中间件函数方法1
	videoGroup := r.Group("/video", m1)
	{
		videoGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/xx"})
		})
		videoGroup.GET("/oo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/oo"})
		})
	}

	// 路由组注册中间件函数方法2
	photoGroup := r.Group("/photo")
	photoGroup.Use(m1, m2)
	{
		photoGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/photo/xx"})
		})
		photoGroup.GET("/oo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/photo/oo"})
		})
	}

	r.GET("/index", indexhandle)

	r.Run(":9090")
}
