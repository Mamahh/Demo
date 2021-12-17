package routers

import (
	"Mamahh/controller"

	"github.com/gin-gonic/gin"
)

//SetupRoute 路由管理
func SetupRoute() *gin.Engine {
	r := gin.Default()

	//解析 HTML 文件
	r.LoadHTMLGlob("./templates/*")

	//规定静态文件的查找路径
	r.Static("/static", "static")

	//demo 路由
	r.GET("/demo", controller.IndexHandler)

	//v1 路由组
	rGroup := r.Group("/v1")
	{
		//添加数据
		rGroup.POST("/todo", controller.CreateATodo)

		//展示数据
		rGroup.GET("/todo", controller.GetTodoList)

		//更新数据
		rGroup.PUT("/todo/:id", controller.UpdateATodo)

		//删除数据
		rGroup.DELETE("/todo/:id", controller.DeleteATodo)

	}
	return r
}
