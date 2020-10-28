package main

import (
	"log"
	"time"
)

func main() {
	c := make(chan int, 1)

	go func() {
		for {
			log.Println("send ready")
			c <- 1
			log.Println("send after")
			//time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for i := range c {
			log.Println("recv:", i)
		}
	}()

	time.Sleep(20 * time.Second)
}
