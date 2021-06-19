package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //Init()
)

var db *sql.DB

func Init_DB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/test_db" //默认端口为3306
	db, err = sql.Open("mysql", dsn)               //用sql.Open方法创建数据库链接池sql.DB 此处不校验密码，只检查dsn的格式是否正确
	if err != nil {
		fmt.Printf("%s 数据格式错误", dsn)
		return
	}
	err = db.Ping() //此处校验密码
	if err != nil {
		fmt.Printf("%s连接失败，请检查密码是否正确\n", dsn)
		return
	}

	return err
}

//定义一个包含数据表属性的类
type user struct {
	runoob_id       int
	runoob_title    string
	runoob_author   string
	submission_date string
}

//根据id查询数据
func query(id int) {
	var u1 user

	//写一条查询单条记录的sql语句
	sql_Str := `select runoob_id,runoob_title,runoob_author,submission_date from runoob_tbl where runoob_id=?;`

	//执行拿到结果,必须使用Scan方法，不然获取不到数据。//该方法会释放数据库链接，从连接池里拿出一个连接查询单条记录
	db.QueryRow(sql_Str, id).Scan(&u1.runoob_id, &u1.runoob_title, &u1.runoob_author, &u1.submission_date)
	fmt.Printf("u1: %#v\n", u1)
}

func main() {
	err := Init_DB() //初始化
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	fmt.Println("数据库连接成功")

	//根据id查询数据
	query(2)
}
