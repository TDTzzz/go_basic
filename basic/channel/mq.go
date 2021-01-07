package main

import (
	"fmt"
)

//消费生产模型
func main() {
	ch := make(chan int, 3)

	nums := 5
	done := make(chan bool, nums)

	for i := 1; i <= nums; i++ {
		go consumer(i, ch, done)
	}
	go producer(ch)

	for i := 1; i <= nums; i++ {
		<-done
	}
}

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(id int, ch chan int, done chan bool) {
	for {
		value, ok := <-ch
		if ok {
			fmt.Printf("id: %d, recv: %d\n", id, value)
		} else {
			fmt.Printf("id: %d, closed\n", id)
			break
		}
	}
	done <- true
}
