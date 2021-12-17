package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB  *gorm.DB
	err error
)

//Todo 定义实例
type Todo struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

//数据库初始化
func initMySQL() (err error) {
	//创建数据库
	//CREATE DATABASE db2;
	dsn := "root:123456@/db2?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	//数据库运行状态查询
	return DB.DB().Ping()

}

func main() {

	err := initMySQL()
	if err != nil {
		panic(DB)
	}

	defer DB.Close()

	//关联实例到数据库
	DB.AutoMigrate(&Todo{})

	r := gin.Default()
	//规定静态文件的查找路径
	r.Static("/static", "static")

	//解析 HTML 文件
	r.LoadHTMLGlob("./templates/*")
	r.GET("/demo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	rGroup := r.Group("/v1")
	{
		//添加数据
		rGroup.POST("/todo", func(c *gin.Context) {
			//前端页面填写待办
			//1. 接受返回数据
			var todo Todo
			c.BindJSON(&todo)
			//2. 将数据存入数据库
			if err = DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})

		//展示数据
		rGroup.GET("/todo", func(c *gin.Context) {
			//1. 刷新数据库的数据
			var todoList []Todo
			if err = DB.Find(&todoList).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})

		//更新数据
		rGroup.PUT("/todo/:id", func(c *gin.Context) {
			//1. 获取要更新的数据id
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
			}
			var todo Todo
			//2. 数据库查询id
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			//绑定结构体(每次前端界面点击待办的按钮会将status的值取反)
			c.BindJSON(&todo)
			//将数据更新到数据库中
			if err = DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})

		//删除数据
		rGroup.DELETE("/todo/:id", func(c *gin.Context) {
			//1. 获取要删除的id
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
			}
			var todo Todo
			//2. 数据库查询id并删除
			if err = DB.Where("id=?", id).Delete(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{id: "deleted"})
			}
		})

	}
	r.Run(":9090")
}
