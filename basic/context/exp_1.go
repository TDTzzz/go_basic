package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	ctx, stop := context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		watchDog(ctx, "【监控狗子1】")
	}()
	go func() {
		defer wg.Done()
		watchDog(ctx, "【监控狗子2】")
	}()
	go func() {
		defer wg.Done()
		watchDog(ctx, "【监控狗子3】")
	}()
	valCtx := context.WithValue(ctx, "userId", 2)
	go func() {
		defer wg.Done()
		getUser(valCtx)
	}()
	time.Sleep(5 * time.Second)
	stop()
	wg.Wait()
}

func watchDog(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			log.Println(name, "停止指令收到，马上停止")
			return
		default:
			log.Println(name, "正在监控...")
		}
		time.Sleep(1 * time.Second)
	}
}

func getUser(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("【获取用户】", "协程退出")
			return
		default:
			userId := ctx.Value("userId")
			log.Println("【获取用户】", "用户ID为", userId)
			time.Sleep(1 * time.Second)
		}
	}
}
