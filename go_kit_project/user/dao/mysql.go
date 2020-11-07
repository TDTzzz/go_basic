package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitMysql(host, port, user, password, dbName string) (err error) {
	//dsn := "root:123456@tcp(127.0.0.1:3306)/test_db?charset=utf8"
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}
	return
}
