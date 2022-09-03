package main

import "fmt"

var done = make(chan bool)
var msg string

func main() {

	go aGo(done)
	//<-done
	done <- true
	fmt.Println(msg)
}

func aGo(done chan bool) {
	msg = "----"
	//done <- true
	<-done
}
