package main

import (
	"fmt"
	"log"
	"reflect"
)

func main() {
	//arr := [3]int{1, 2, 3}
	//log.Println(arr[1:2]) //数组下标创建切片
	//testSlice()

	//testAppend()
	testAppend2()
	//testAppend3()
	//testAppend4()
}

//3种创建方法

//字面量的形式创建slice
func testSlice() {
	var vstat [3]int
	slice := vstat[1:]
	fmt.Println(reflect.TypeOf(vstat))
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(slice)
}

//扩容问题
func testAppend() {
	s := make([]int, 0)
	oldCap := cap(s)

	for i := 0; i < 2048; i++ {
		s = append(s, i)
		newCap := cap(s)

		if oldCap != newCap {
			fmt.Printf("[%d -> %4d] cap = %-4d  |  after append %-4d cap =%-4d\n", 0, i-1, oldCap, i, newCap)
			oldCap = newCap
		}
	}
}

//扩容问题2
func testAppend2() {
	s := []int{1, 2}
	fmt.Printf("len=%d,cap=%d", len(s), cap(s))
	s = append(s, 3, 4, 5)
	//s = append(s, 3)
	//s = append(s, 4)
	//s = append(s, 5)
	fmt.Printf("len=%d,cap=%d", len(s), cap(s))
}

//扩容问题3
func testAppend3() {
	//s := []int{}
	//var s []int
	s2 := make([]int, 0)
	s := make([]int, 0)
	s = nil
	log.Println(s, len(s), cap(s), s == nil, reflect.TypeOf(s), reflect.TypeOf(s2))
	s = append(s, 1)
	log.Println(s)
}

//empty slice/nil slice
func testAppend4() {
	var nilSlice []string
	emptySlice0 := make([]int, 0)
	var emptySlice1 = []string{}

	fmt.Printf("\nNil:%v Len:%d Capacity:%d", nilSlice == nil, len(nilSlice), cap(nilSlice))
	fmt.Printf("\nnil:%v Len:%d Capacity:%d", emptySlice0 == nil, len(emptySlice0), cap(emptySlice0))
	fmt.Printf("\nnil:%v Len:%d Capacity:%d", emptySlice1 == nil, len(emptySlice1), cap(emptySlice1))
}
