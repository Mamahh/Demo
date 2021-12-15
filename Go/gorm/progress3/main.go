package main

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//User 1.1 定义类型
// type User struct {
// 	ID   int64
// 	Name string `gorm:"default:'小王子'"` //默认值为 小王子
// 	Age  int64
// }

//User 1.2 定义类型
// type User struct {
// 	ID   int64
// 	Name *string `gorm:"default:'小王子'"` //默认值为 小王子
// 	Age  int64
// }

//User 1.3 定义类型
type User struct {
	ID   int64
	Name sql.NullString `gorm:"default:'小王子'"` //默认值为 小王子
	Age  int64
}

//Duser 定义类型
type Duser struct {
	gorm.Model        //ID CreateAt UpdateAt DeleteAt
	Name       string `gorm:"default:'小王子'"` //默认值为 小王子
	Age        int64
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(db)
	}
	defer db.Close()
	//2. 把模型和数据库的表关联
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Duser{})

	//3. 创建
	// u := User{Name:"", Age: 18} // 1.1默认值小王子
	// u := User{Name: new(string), Age: 28}// 1.2强行置空
	u := User{Name: sql.NullString{"", true}, Age: 38} // 1.3强行置空
	u1 := Duser{Name: "玛卡巴卡", Age: 9}

	fmt.Println(db.NewRecord(&u)) //判断主键是否为空
	db.Debug().Create(&u)         //添加Debug
	db.Debug().Create(&u1)        //添加Debug
	fmt.Println(db.NewRecord(&u)) //判断主键是否为空
}
