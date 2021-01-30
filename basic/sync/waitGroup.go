package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Done()
	wg.Wait()
	log.Println("end!!!")
}
