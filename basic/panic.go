package main

import (
	"fmt"
	"os"
	"time"
)

//panic 只会触发当前 Goroutine 的 defer；
//recover 只有在 defer 中调用才会生效；
//panic 允许在 defer 中嵌套多次调用；

func main() {
	//defer println("in main")
	//go func() {
	//	defer println("in goroutine")
	//	panic("")
	//}()
	//
	//time.Sleep(1 * time.Second)

	testRecover()
}

func testPanic() {
	defer fmt.Println("in main")
	defer func() {
		defer func() {
			panic("panic again and again")
		}()
		panic("panic again")
	}()

	panic("panic once")
}

func testRecover() {
	defer fmt.Println("defer main")
	var user = os.Getenv("USER_")
	go func() {
		defer func() {
			fmt.Println("defer caller")
			if err := recover(); err != nil {
				fmt.Println("recover success.err:", err)
			}
		}()
		func() {
			defer func() {
				fmt.Println("defer here")
			}()

			if user == "" {
				panic("should set user env.")
			}
			fmt.Println("after panic")
		}()
	}()

	time.Sleep(1 * time.Second)
	//fmt.Printf("get result %d\r\n", result)
}
