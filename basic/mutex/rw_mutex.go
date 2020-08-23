package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	var rwMutex sync.RWMutex

	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func(i int) {
			rwMutex.RLock()
			defer rwMutex.RUnlock()
			fmt.Println("Read data ", i)
			wg.Done()
			time.Sleep(2 * time.Second)
		}(i)

		go func(i int) {
			rwMutex.Lock()
			defer rwMutex.Unlock()
			fmt.Println("Write data ", i)
			wg.Done()
			time.Sleep(2 * time.Second)
		}(i)
	}
	wg.Wait()
}
