package dialect

import "reflect"

//抽象出各个数据库差异的部分

var dialectsMap = map[string]Dialect{}

type Dialect interface {
	DataTypeOf(typ reflect.Value) string //将go的数据类型转化成该数据库的类型
	TableExistSQL(tableName string) (string, []interface{})
}

func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}
