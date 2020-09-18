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
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//var user User
	var users []User

	db.Table("user").Find(&users, []int{1, 2})
	log.Println(users)
}
