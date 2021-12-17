package main

import (
	"Mamahh/dao"
	"Mamahh/model"
	"Mamahh/routers"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	//初始化数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.DB.Close()

	//关联实例到数据库
	dao.DB.AutoMigrate(&model.Todo{})

	//注册路由
	r := routers.SetupRoute()

	//启动
	r.Run(":9090")
}
