package session

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go_basic/7days/gee-orm/day2-reflect-schema/dialect"
	"os"
	"testing"
)

var (
	TestDB      *sql.DB
	TestDial, _ = dialect.GetDialect("mysql")
)

func TestMain(m *testing.M) {
	source := "root:123456@tcp(127.0.0.1:3306)/test_db?charset=utf8"
	TestDB, _ = sql.Open("mysql", source)
	code := m.Run()
	_ = TestDB.Close()
	os.Exit(code)
}

func NewSession() *Session {
	return New(TestDB, TestDial)
}

func TestSession_Exec(t *testing.T) {
	s := NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS users;").Exec()
	_, _ = s.Raw("CREATE TABLE users(" +
		"`id` INT UNSIGNED AUTO_INCREMENT," +
		"name VARCHAR(255) NOT NULL DEFAULT ''," +
		"age TINYINT(4) NOT NULL DEFAULT 0," +
		"PRIMARY KEY(`id`)" +
		")ENGINE=InnoDB DEFAULT CHARSET=utf8;").Exec()
	result, _ := s.Raw("INSERT INTO users(`name`,`age`) values (?,?)", "TDTzzz", 23).Exec()
	if count, err := result.RowsAffected(); err != nil || count != 1 {
		t.Fatal("expect 1,but got", count)
	}
}

//func TestSession_QueryRows(t *testing.T) {
//	s := NewSession()
//	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
//	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
//	row := s.Raw("SELECT count(*) FROM User").QueryRow()
//	var count int
//	if err := row.Scan(&count); err != nil || count != 0 {
//		t.Fatal("failed to query db", err)
//	}
//}
