package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test_db?charset=utf8"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//var user User
	var users []User

	db.Table("users").Find(&users, []int{1, 2})
	log.Println(users)
}
