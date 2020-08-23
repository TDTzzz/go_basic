package main

import (
	"log"
	"reflect"
)

func main() {
	var a int = 50
	v := reflect.ValueOf(a)
	t := reflect.TypeOf(a)
	k := reflect.Kind(a)

	log.Println(v, t, k)
}
