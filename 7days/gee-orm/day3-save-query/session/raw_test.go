package session

import (
	"database/sql"
	"go_basic/7days/gee-orm/day3-save-query/dialect"
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
