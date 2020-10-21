package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

var db *sql.DB

//func init() {
//	db, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:13306)/fileserver?charset=utf8")
//
//	db.SetMaxOpenConns(1000)
//	err := db.Ping()
//
//	if err != nil {
//		fmt.Println("Failed to connect to mysql,err:" + err.Error())
//		os.Exit(1)
//	}
//}

//数据库配置
//Db数据库连接池
var DB *sql.DB

//注意方法名大写，就是public
func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	//DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	//DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		log.Println(err)
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connect success")
}

func main() {
	InitDB()
	//InsertData(DB)
	//SelectData(DB)
	//
	SelectManyRow(DB)

}

func CreateTable(DB *sql.DB) {

}

func InsertData(DB *sql.DB) {
	res, err := DB.Exec("insert INTO user(name,age) values (?,?)", "HAHAHA", 25)
	if err != nil {
		fmt.Printf("Insert data failed,err:%v", err)
		return
	}
	lastInsertID, err := res.LastInsertId()

	rowsaffected, err := res.RowsAffected()

	log.Println(lastInsertID)
	log.Println(rowsaffected)

}

func SelectData(DB *sql.DB) {

	user := new(User)
	row := DB.QueryRow("select * from user")

	row.Scan(&user.Id, &user.Name, &user.Age)
	log.Println(user)
}

func SelectManyRow(DB *sql.DB) {

	user := new(User)
	//res := []User{}
	row, _ := DB.Query("select * from user")

	for row.Next() {
		row.Scan(&user.Id, &user.Name, &user.Age)
		log.Println(user)
	}
}
