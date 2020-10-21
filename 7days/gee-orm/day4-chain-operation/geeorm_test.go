package day4_chain_operation

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func OpenDB(t *testing.T) *Engine {
	t.Helper()
	source := "root:123456@tcp(127.0.0.1:3306)/test_db?charset=utf8"
	engine, err := NewEngine("mysql", source)
	if err != nil {
		t.Fatal("failed to connect", err)
	}
	return engine
}

func TestNewEngine(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
}
