package dialect

import (
	"log"
	"reflect"
)

type mysql struct{}

func init() {
	log.Println("注册mysql-dialect!!!")
	RegisterDialect("mysql", &mysql{})
}

func (m mysql) DataTypeOf(typ reflect.Value) string {
	panic("implement me")
}

func (m mysql) TableExistSQL(tableName string) (string, []interface{}) {
	args := []interface{}{tableName}
	return "SELECT table_name FROM information_schema.TABLES WHERE table_name= ?", args
}
