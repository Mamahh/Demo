package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	//DB 全局 gorm
	DB  *gorm.DB
	err error
)

//InitMySQL 数据库初始化
func InitMySQL() (err error) {
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
