package main

import "log"

//defer相关
func main() {
	log.Println("return:", testDefer())
	log.Println("return:", testDefer2())
}

//这里的return 会是0
func testDefer() int {
	var i int
	defer func() {
		i++
		log.Println("defer1", i)
	}()
	defer func() {
		i++
		log.Println("defer2", i)
	}()
	return i
}

func testDefer2() (i int) {
	defer func() {
		i++
		log.Println("defer1", i)
	}()
	defer func() {
		i++
		log.Println("defer2", i)
	}()
	return i
}
