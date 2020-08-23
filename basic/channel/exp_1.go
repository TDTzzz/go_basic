package main

import (
	"log"
	"time"
)

//阻塞雨非阻塞
func main() {
	c := make(chan int)
	go send(c)
	go recv(c)
	time.Sleep(3 * time.Second)
	close(c)
}

func send(c chan int) {
	for i := 0; i < 10; i++ {
		log.Println("send before:", i)
		c <- i
		log.Println("send after:", i)
	}
}

func recv(c chan int) {
	for i := range c {
		log.Println("recv:", i)
	}
}
