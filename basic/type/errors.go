package main

import (
	"fmt"
)

func div(a, b int) {

	defer func() {
		//recover() 这个内建函数被用于从异常或错误场景中恢复：让程序可以从 panicking 重新获得控制权，停止终止过程进而恢复正常执行。
		if r := recover(); r != nil {
			fmt.Printf("捕获到异常：%s\n", r)
		}
	}()

	if b < 0 {

		panic("除数需要大于0")
	}

	fmt.Println("余数为：", a/b)

}

func main() {
	// 捕捉内部的异常
	div(10, 0)

	// 捕捉主动的异常
	div(10, -1)
}

//defer：先进后出
func deferCase() {
	//输出:  hello  world  !!!
	defer fmt.Print(" !!! ")
	defer fmt.Print(" world ")
	fmt.Print(" hello ")
}
