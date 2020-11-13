package main

import "fmt"

//明明2个map的指针不一样，但是却可以修改Map的内容，magic magic
//好文章：https://www.flysnow.org/2018/02/24/golang-function-parameters-passed-by-value.html
func main() {
	persons := make(map[string]int)
	persons["张三"] = 19

	mp := &persons

	fmt.Printf("原始map的内存地址是：%p\n", mp)
	testMap(persons)
	fmt.Println("map值被修改了，新值为:", persons)
}

func testMap(p map[string]int) {
	fmt.Printf("函数里接收到map的内存地址是：%p\n", &p)
	p["张三"] = 20
}
