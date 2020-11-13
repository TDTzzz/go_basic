package main

import (
	"fmt"
	"log"
)

//好文章:https://www.flysnow.org/2018/02/24/golang-function-parameters-passed-by-value.html
//Go里函数都是值传递，无引用传递
func main() {
	i := 10
	ip := &i
	log.Println(*&ip)
	log.Println(ip)
	fmt.Printf("原始指针的内存地址是：%p\n", &ip)
	modify(ip)
	fmt.Println("int值被修改了，新值为:", i)
}

func modify(ip *int) {
	fmt.Printf("函数里接收到的指针的内存地址是：%p\n", &ip)
	*ip = 1
}
