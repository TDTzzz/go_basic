package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	exp1()
}

//超时控制
func exp1() {
	ctx, cancle := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancle()
	go handle(ctx, 1500*time.Millisecond)

	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("ctx超时:", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with:", duration)
	}
}
