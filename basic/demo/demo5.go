package main

import (
	"log"
	"sync"
)

func main() {
	ch := make(chan int)

	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for i := range ch {
			log.Println(i)
		}
	}()

	wg.Wait()
	//time.Sleep(2 * time.Second)
}
