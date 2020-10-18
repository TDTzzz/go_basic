package session

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go_basic/7days/gee-orm/day4-chain-operation/dialect"
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

func TestSession_QueryRow(t *testing.T) {
	s := NewSession()
	row := s.Raw("select count(*) from users").QueryRow()
	var count int
	if err := row.Scan(&count); err != nil {
		t.Fatal("failed to query db", err, count)
	}
}
