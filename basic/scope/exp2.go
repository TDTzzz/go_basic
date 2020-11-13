package main

import "log"

func main() {
	p := 1
	log.Println(p, &p)
	test(p)
	log.Println(p, &p)
}

func test(p interface{}) {
	p = 2
	log.Println(p, &p)
}
