package main

import (
	"fmt"
	"log"
)

//var map1 map[keytype]valuetype
//key可以是任意可以用==或者!=操作符比较的类型

//map是引用类型
//map容量:和数组不同，map 可以根据新增的 key-value 对动态的伸缩，因此它不存在固定长度或者最大限制。
//但是你也可以选择标明 map 的初始容量 capacity，就像这样：make(map[keytype]valuetype，cap)。
func main() {
	//case1()

	//var m map[string]int
	m := make(map[string]int)
	//m := map[string]int{}

	log.Println(m == nil)

	if _, ok := m["two"]; !ok {
		fmt.Println("no entry")
	}

	m["tt"] = 11
	delete(m, "tt")
	log.Println(m)

}

//对一个nil的slice添加元素没问题
//但是对nil的map添加会报panic
func case1() {
	var m map[string]int
	m["one"] = 1
}

//map的range
//
func case2() {
	data := []int{1, 2, 3}
	for _, v := range data {
		//这里的v改动不会改变data
		v *= 10
	}
	log.Println(data)

	for i, _ := range data {
		//应该直接修改data
		data[i] *= 10
	}

	log.Println(data)
}
