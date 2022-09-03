package main

import (
	"fmt"
	"runtime"
	"time"
)

//如何超时取消协程

func main() {

	for i := 0; i < 1000; i++ {
		timeout(doGood)
	}

	fmt.Println(runtime.NumGoroutine())
	time.Sleep(time.Second)
	fmt.Println(runtime.NumGoroutine())
}

func timeout(f func(chan bool)) error {

	//这里加不加缓冲区区别很大，有了缓冲区，doSomething 1s后向done里发送后
	//协程可以直接退出
	done := make(chan bool, 1)

	//不加缓冲区，超时提前返回，doSomething 1s后向done里发送由于没有接受方，
	//导致阻塞，协程无法退出
	//done := make(chan bool)
	go f(done)

	select {
	case <-done:
		fmt.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}

}

func doSomething(done chan bool) {
	time.Sleep(time.Second)
	//如果done没接收者了且没缓冲，则会阻塞
	done <- true
}

//设置缓冲区还有一种方式，用select
func doGood(done chan bool) {
	time.Sleep(time.Second)
	select {
	case done <- true:
	default:
		return
	}
}
