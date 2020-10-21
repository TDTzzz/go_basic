package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//Mysql的事务

func main() {
	source := "root:123456@tcp(127.0.0.1:3306)/test_db?charset=utf8"
	db, _ := sql.Open("mysql", source)

	//1.开启事务
	tx, _ := db.Begin()
	_, err1 := tx.Exec("insert into users(`name`,`age`) values (?,?)", "ddd", 21)
	_, err2 := tx.Exec("insert into users(`name`,`age`) values (?,?)", "ssss", 22)

	if err1 != nil || err2 != nil {
		_ = tx.Rollback()
		log.Println("Rollback", err1, err2)
	} else {
		_ = tx.Commit()
		log.Println("Commit")
	}

}
