package main

import "fmt"

//闭包
//函数+引用环境=闭包
func main() {
	accAdd := Accumulate(1)
	fmt.Println(accAdd())
	fmt.Println(accAdd())
}

func Accumulate(value int) func() int {
	return func() int {
		value++
		return value
	}
}
