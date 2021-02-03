package main

import (
	"fmt"
	"log"
	"reflect"
)

func main() {
	var i int = 3
	iv := reflect.ValueOf(i)
	it := reflect.TypeOf(i)
	log.Println(iv, it)
	fmt.Println(iv.Interface().(int))

	test()
}

//通过反射修改变量
func test() {
	i := 3
	ipv := reflect.ValueOf(&i)
	log.Println(ipv.Kind())
	ipv.Elem().SetInt(4)
	fmt.Println(i)
}


