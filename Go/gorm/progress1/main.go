package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//UserInfo userinfo
type UserInfo struct {
	Name  string
	Age   string
	Hobby string
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(db)
	}
	defer db.Close()
	u := UserInfo{
		Name:  "Mamahh",
		Age:   "7",
		Hobby: "playgame",
	}

	//创建表 自动迁移 （把结构体和数据表进行对应）
	db.AutoMigrate(&UserInfo{})

	//创建C
	// db.Create(&u)

	//查询表中第一条数据保存到u
	db.First(&u)
	fmt.Printf("DB : %#v\n", u)

	//更新
	db.Model(&u).Update("Hobby", "王者荣耀")

	//删除
	db.Delete(&u)
}
