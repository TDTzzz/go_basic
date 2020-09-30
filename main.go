package main

import (
	"fmt"
)

func main() {
	var data int
	go func() {
		data++
	}()
	//time.Sleep(time.Millisecond)
	if data == 0 {
		fmt.Println(data)
	}
}
