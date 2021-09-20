package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//关系型数据库 ORM(Object Relational Mapping)对象关系映射
//mysql中字段和go中结构体来映射

type User struct {
	ID       int `gorm:"primaryKey"`
	Username string
	Password string
}

func SqlBygo() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:42128340HDC@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	var user []User
	db.Table("db_table").Find(&user)

	fmt.Printf("%#v", user[0:5])

}

func main() {
	SqlBygo()
}
