package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {

	letter, number := make(chan bool), make(chan bool)
	wg := sync.WaitGroup{}

	//打印数字
	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				letter <- true
				break
			default:
				break
			}
		}
	}()
	wg.Add(1)

	//打印字母
	go func(wg *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0

		for {
			select {
			case <-letter:
				if i >= strings.Count(str, "")-1 {
					wg.Done()
					return
				}
				fmt.Print(str[i : i+1])
				i++
				//if i >= strings.Count(str, "") {
				//	i = 0
				//}
				fmt.Print(str[i : i+1])
				i++
				number <- true
				break
			default:
				break
			}
		}
	}(&wg)
	number <- true
	wg.Wait()
}
