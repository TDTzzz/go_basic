package main

import (
	"log"
)

func main() {
	ch := make(chan int)
	go producer(ch)
	go consumer(ch)

	//time.Sleep(2 * time.Second)
}

func consumer(ch chan int) {
	for {
		value, ok := <-ch
		if ok {
			log.Println(value)
		} else {
			//close了 关闭
			log.Println("close-----")
			break
		}
	}
}

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
