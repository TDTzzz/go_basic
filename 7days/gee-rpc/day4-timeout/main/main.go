package main

import (
	day4_timeout "go_basic/7days/gee-rpc/day4-timeout"
	"log"
	"net"
	"sync"
	"time"
)

type Foo int

func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

type Args struct{ Num1, Num2 int }

func startServer(addr chan string) {
	//1.注册服务
	var foo Foo
	if err := day4_timeout.Register(&foo); err != nil {
		log.Fatal("register error:", err)
	}
	//2.建立通信
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on ", l.Addr())
	addr <- l.Addr().String()
	//3.开启server的accept
	day4_timeout.Accept(l)
}

func main() {
	addr := make(chan string)
	go startServer(addr)
	//Client建立连接
	client, _ := day4_timeout.Dial("tcp", <-addr)
	defer func() {
		_ = client.Close()
	}()

	time.Sleep(3 * time.Second)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			args := &Args{Num1: i, Num2: i * i}
			var reply int
			if err := client.Call("Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum error:", err)
			}
			log.Printf("%d + %d = %d", args.Num1, args.Num2, reply)
		}()
	}
}
