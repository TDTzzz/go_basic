package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan int)

	go rec(ch)
	go prod(ch)
	time.Sleep(time.Second)
}

func rec(ch chan int) {
	ch <- 1
}

func prod(ch chan int) {
	if v, ok := <-ch; ok {
		log.Println("open value:", v)
	} else {
		log.Println("close")
	}

	close(ch)
	if v, ok := <-ch; ok {
		log.Println("open value:", v)
	} else {
		log.Println("close")
	}
}
