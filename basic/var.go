package main

import (
	"fmt"
	"log"
)

//内存分配
func main() {
	//下面这个会报错，因为指针类型声明了但是没分配内部，无法赋值
	var sp *string
	*sp = "飞雪无情"
	fmt.Println(*sp) //panic: runtime error: invalid memory address or nil pointer dereference

	//对于值类型来讲，即使声明一个对象，没对其初始化，该变量也会有分配好哦的内存
	var sp2 string
	log.Println(&sp2)
	sp2 = "阿巴阿巴"
	fmt.Println(sp2)

	//
	test()
}

func test() {
	var sp *string
	sp = new(string)
	*sp = "阿巴阿巴"
	fmt.Println(*sp)
}

//func makeTest() {
//	m := make(map[string]int, 10)
//}
