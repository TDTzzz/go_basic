package main

import (
	"fmt"
	"unsafe"
)

//切片的函数传递问题
func main() {
	//s := []int{1, 2, 3}
	//fmt.Println(unsafe.Pointer(&s))
	//f(s)
	//fmt.Println(unsafe.Pointer(&s))
	//fmt.Println(s)

	//test2()
	test4()
}

func f(s []int) {
	fmt.Println(unsafe.Pointer(&s))
	for i := range s {
		s[i]++
	}
}

func test2() {

	s := []int{1, 2, 3}
	newS := myAppend(s)

	fmt.Println(s)
	fmt.Println(newS)

	myAppendPtr(&s)
	fmt.Println(s)
}

func myAppend(s []int) []int {
	s = append(s, 100)
	return s
}

func myAppendPtr(s *[]int) {
	*s = append(*s, 100)
}

func test4() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	s2 := s1[2:6:7]

	s2 = append(s2, 100)
	s2 = append(s2, 200)

	s1[2] = 20

	fmt.Println(s1, cap(s1))
	fmt.Println(s2, cap(s2))
	fmt.Println(slice, cap(slice))
}
