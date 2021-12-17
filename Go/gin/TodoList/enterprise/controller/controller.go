package controller

import (
	"Mamahh/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
 url     --> controller  --> logic   -->    model
请求来了  -->  控制器      --> 业务逻辑  --> 模型层的增删改查
*/

//IndexHandler html渲染
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

//CreateATodo 新增Todo清单
func CreateATodo(c *gin.Context) {
	//前端页面填写待办
	//1. 接受返回数据
	var todo model.Todo
	c.BindJSON(&todo)
	//2. 将数据存入数据库
	if err := model.CreateATodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

//GetTodoList 获取Todo数据
func GetTodoList(c *gin.Context) {
	//1. 刷新数据库的数据
	todoList, err := model.GetAllToDo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

//UpdateATodo 更新Todo的一条数据 by id
func UpdateATodo(c *gin.Context) {
	//1. 获取要更新的数据id
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
	}
	//2. 数据库查询id
	todo, err := model.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	//绑定结构体(每次前端界面点击待办的按钮会将status的值取反)
	c.BindJSON(&todo)
	//将数据更新到数据库中
	if err := model.UpdateATodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

//DeleteATodo by id
func DeleteATodo(c *gin.Context) {
	//1. 获取要删除的id
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
	}
	//2. 数据库查询id并删除
	err := model.DeleteATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
