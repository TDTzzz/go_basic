package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	//once()
	syncMap()
}

func once() {
	var once sync.Once

	onced := func() {
		fmt.Println("onced")
	}
	onces := func() {
		fmt.Println("onces")
	}
	for i := 0; i < 10; i++ {
		once.Do(onces)
		go once.Do(onced)
	}
	time.Sleep(1 * time.Second)
}

func waitGroup() {

}

//普通的map不支持线程安全
func syncMap() {
	var m sync.Map

	m.Store("name", "Jack")
	m.Store("age", 21)

	v, ok := m.LoadOrStore("name1", "aaaa")
	log.Println(v, ok)

	f := func(k, v interface{}) bool {
		log.Println(k, "-", v)
		return true
	}

	m.Range(f)
}
