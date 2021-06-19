# Go 操作mariadb

## mariadb配置

### 安装

```shell
sudo apt install mariadb-server mariadb-client -y
```

### 初始化

```shell
sudo mysql_secure_installation
```

配置完密码，一路回车即可。

### 解决远程无法连接的问题

当设置了无密码登录时，远程访问即使密码输入正确也是无法连接的，需要移出无密码登录的插件。

判断系统是否含有无密码的插件命令

```mariadb
MariaDB [(none)]> select user, plugin from mysql.user;
```

```mysql
1 正常mysql
2 mysql> select user, plugin from mysql.user where plugin = 'mysql_native_password';
3 +-----------+-----------------------+
4 | user      | plugin                |
5 +-----------+-----------------------+
6 | root      | mysql_native_password |
7 +-----------+-----------------------+
8 8 rows in set (0.00 sec)
```

```mariadb
1 不正常的
2 
3 MariaDB [(none)]> select user, plugin from mysql.user;
4 +------+-------------+
5 | user | plugin      |
6 +------+-------------+
7 | root | unix_socket |
8 +------+-------------+
9 1 row in set (0.00 sec)
```

看到这里应该发现问题了，按照正常的修改就行了

如下：

```shell
 1 sudo systemctl stop mariadb.service
 2 sudo mysqld_safe --skip-grant-tables
 3 进去mysql执行如下命令：
 4 MariaDB [(none)]> UPDATE mysql.user SET authentication_string = PASSWORD('mypassword'), plugin = 'mysql_native_password' WHERE User = 'root' AND Host = 'localhost';
 5 MariaDB [(none)]> FLUSH PRIVILEGES;
 6 验证：
 7 MariaDB [(none)]> select user, plugin from mysql.user
 8     -> ;
 9 +------+-----------------------+
10 | user | plugin                |
11 +------+-----------------------+
12 | root | mysql_native_password |
13 +------+-----------------------+
14 1 row in set (0.01 sec)
15 
16 先杀死mysql sudo kill -9 pid
17 启动：
18 sudo systemctl start mariadb.service
```

最后验证下：需要密码了

```shell
~ mysql
ERROR 1045 (28000): Access denied for user 'root'@'localhost' (using password: NO)
```

此时远程就可以通过密码连接了。

### 创建数据库demo

#### 登陆数据库

执行命令: `mysql -u root -p` 登陆到数据库

Enter password:******* （输入root用户密码即可）

#### 数据库基本命令的使用

执行命令:`create database mydatabase;`( 创建数据库，mydatabase表示数据库名称，可以自行命名)

执行命令：`mysqladmin -u root -p create mydatabase`（不进入数据库直接在终端创建数据库）



show databases; （展示所有的数据库），返回值如下：

```mariadb
MariaDB [(none)]> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| test_db            |
+--------------------+
4 rows in set (0.000 sec)
```

```mariadb
use mydatabase; (使用名称为mydatabase的数据库，可以在这个数据库里面创建表)，返回值如下：

Database changed
```

show tables; （展示所有的表）

创建名称为没有runoob_tbl这个表就创建一个名称为runoob_tbl的表，并且设置表头（表头一定需要，否则会报错）

```
create table if not exists `runoob_tbl`(

  `runoob_id` INT UNSIGNED AUTO_INCREMENT,

  `runoob_title` VARCHAR(100) NOT NULL,

  `runoob_author` VARCHAR(40) NOT NULL,

  `submission_date` DATE,

  PRIMARY KEY ( `runoob_id` ))ENGINE=InnoDB DEFAULT CHARSET=utf8; （设置主键，便于与副表连接）

返回值如下：Query OK, 0 rows affected (0.01 sec)


```

上述实例解析：

- 如果你不想字段为 NULL 可以设置字段的属性为 NOT NULL， 在操作数据库时如果输入该字段的数据为NULL ，就会报错。
- AUTO_INCREMENT定义列为自增的属性，一般用于主键，数值会自动加1。
- PRIMARY KEY关键字用于定义列为主键。 您可以使用多列来定义主键，列间以逗号分隔。
- ENGINE 设置存储引擎，CHARSET 设置编码



插入数据到表中

```mariadb
insert into runoob_tbl

   (runoob_title, runoob_author, submission_date)

  values

   ("学习 PHP", "菜鸟教程", NOW());

返回值如下：Query OK, 1 rows affected, 1 warnings (0.01 sec)


select * from runoob_tbl;查询runoob_tbl表中的数据，返回值如下:

mysql> select * from runoob_tbl;

+-----------+--------------+---------------+-----------------+

| runoob_id | runoob_title | runoob_author | submission_date |

+-----------+--------------+---------------+-----------------+

|         1 | 学习 PHP     | 菜鸟教程      | 2020-08-25      |


update runoob_tbl set runoob_title='apache' where runoob_title="学习 PHP"; （学习PHP  修改为apache，返回值）

Query OK, 2 rows affected (0.01 sec)

Rows matched: 1  Changed: 1  Warnings: 0

输入命令:exit (退出数据库)
```

最终的表格demo如下：

```mariadb
MariaDB [test_db]> show tables;
+-------------------+
| Tables_in_test_db |
+-------------------+
| runoob_tbl        |
+-------------------+
1 row in set (0.000 sec)

MariaDB [test_db]> select * from runoob_tbl;
+-----------+--------------+---------------+-----------------+
| runoob_id | runoob_title | runoob_author | submission_date |
+-----------+--------------+---------------+-----------------+
|         1 | 学习 PHP     | 菜鸟教程      | 2021-06-19      |
|         2 | Go           | Bilibili      | 2021-06-19      |
+-----------+--------------+---------------+-----------------+
2 rows in set (0.000 sec)
```

至此，我们可以开始用Go去操作数据库了。

## Go mysql环境配置

### 配置Go module

```shell
go mod init github.com
```

### 导入go-sql-driver类库

```shell
go get -u github.com/go-sql-driver/mysql
```

### 具体使用见代码注释：

```go
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //Init()
)

var db *sql.DB

func Init_DB() (err error) {
	dsn := "user:password@tcp(127.0.0.1:3306)/test_db" //默认端口为3306
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
```

输出：

```shell
$ go run mysql.go
数据库连接成功
u1: main.user{runoob_id:2, runoob_title:"Go", runoob_author:"Bilibili", submission_date:"2021-06-19"}
```

