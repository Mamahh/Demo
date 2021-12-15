package main

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//User 表名默认为结构体名的复数 users
type User struct {
	gorm.Model   //内嵌gorm.Model //默认ID 为primary key
	Name         string
	Age          sql.NullInt64 //零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type=varchar(120);unique_index"`
	Role         string  `gorm:"size:255"`        //设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` //设置会员号唯一且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  //设置num为自增类型
	Address      string  `gorm:"index:addr"`      //给address字段创建addr的索引
	IgnoreMe     string  `gorm:"-"`               //忽略本字段

}

//TableName 重命名User 表名为 Mamahh
func (User) TableName() string {
	return "Mamahh"
}

//Animal x
type Animal struct {
	//自定义AnimalID 为主键
	AnimalID int64 `gorm:"primary_key"`
	Name     string
	Age      int64 `gorm:"column:animal_age"`
}

func main() {
	//修改默认的表名规则
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "prefix_" + defaultTableName
	}
	db, err := gorm.Open("mysql", "root:123456@/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(db)
	}
	defer db.Close()

	//创建表时禁止默认为复数名
	db.SingularTable(true)

	//自动匹配结构体
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Animal{})

	//创建表并起名
	// db.Table("My_table").CreateTable(&User{})
}
