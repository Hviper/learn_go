package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//关系型数据库 ORM(Object Relational Mapping)对象关系映射
//mysql中字段和go中结构体来映射

func SqlBygo() {
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	rows, e := db.Query("select * from db_table")
	if e != nil {
		log.Fatal(e)
	}
	fmt.Printf("%#v", rows)
	defer db.Close()
}

func main() {
	SqlBygo()
}
