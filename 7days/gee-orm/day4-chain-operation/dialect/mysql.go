package dialect

import (
	"fmt"
	"reflect"
)

type mysql struct{}

//todo 不知道这个是干嘛的
var _ Dialect = (*mysql)(nil)

func init() {
	RegisterDialect("mysql", &mysql{})
}

//func (m mysql) DataTypeOf(typ reflect.Value) string {
//	switch typ.Kind() {
//	case reflect.Bool:
//		return "bool"
//	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
//		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
//		return "integer"
//	case reflect.Int64, reflect.Uint64:
//		return "bigint"
//	case reflect.Float32, reflect.Float64:
//		return "real"
//	case reflect.String:
//		return "text"
//	case reflect.Array, reflect.Slice:
//		return "blob"
//	case reflect.Struct:
//		if _, ok := typ.Interface().(time.Time); ok {
//			return "datetime"
//		}
//	}
//	panic(fmt.Sprintf("invalid sql type %s (%s)", typ.Type().Name(), typ.Kind()))
//}

func (m mysql) TableExistSQL(tableName string) (string, []interface{}) {
	args := []interface{}{tableName}
	return "SELECT table_name FROM information_schema.TABLES WHERE table_name= ?", args
}

func (m mysql) DataTypeOf(typ reflect.Value) string {
	switch typ.Kind() {
	case reflect.Bool:
		return "boolean"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
		return "int"
	case reflect.Int64, reflect.Uint64:
		return "bigint"
	case reflect.Float64, reflect.Float32:
		//忽略了double 和 decimal的差异
		return "float"
	case reflect.String:
		//测试先都按255来 忽略text的类型
		return "varchar(255)"
	}
	panic(fmt.Sprintf("invalid sql type %s (%s)", typ.Type().Name(), typ.Kind()))
}
