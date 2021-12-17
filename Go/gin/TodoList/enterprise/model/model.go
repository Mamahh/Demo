package model

import (
	"Mamahh/dao"
)

//Todo 定义实例
type Todo struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

//CreateATodo 新增todo数据
func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(todo).Error
	return
}

// GetAllToDo 获取所有数据
func GetAllToDo() (todoList []*Todo, err error) {
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

//GetATodo 查询一条数据 by id
func GetATodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = dao.DB.Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

//UpdateATodo 更新一条数据
func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

//DeleteATodo 删除一条数据
func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
