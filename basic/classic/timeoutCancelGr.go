package main

import (
	"context"
	"fmt"
	"time"
)

//超时取消协程
func main() {

	//methodOne()
	//methodTwo()
	methodThree()
}

//方式一 ctx+time.After
func methodOne() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*800))
	defer cancel()

	go func(ctx context.Context) {
		time.Sleep(time.Second)
		fmt.Println("-----")
	}(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("call success")
	case <-time.After(time.Duration(time.Millisecond * 900)):
		fmt.Println("timeout!!!")
	}

}

//方式二 ctx+time.NewTimer
func methodTwo() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*800))
	defer cancel()
	timer := time.NewTimer(time.Duration(time.Millisecond * 900))

	go func(ctx context.Context) {

	}(ctx)

	select {
	case <-ctx.Done():
		timer.Stop()
		timer.Reset(time.Second)
		fmt.Println("call success")
	case <-timer.C:
		fmt.Println("timeout!!!")
	}
}

//方式三 select+缓冲channel

func methodThree() {

	done := make(chan struct{}, 1)

	go func(done chan struct{}) {
		//逻辑耗时
		time.Sleep(time.Second)
		done <- struct{}{}
	}(done)

	select {
	case <-done:
		fmt.Println("call success")
	case <-time.After(time.Millisecond * 1100):
		fmt.Println("超时")
	}
}
