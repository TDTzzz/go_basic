package main

import (
	"fmt"
	"log"
	"sort"
)

//var map1 map[keytype]valuetype
//key可以是任意可以用==或者!=操作符比较的类型

//1.map绝不用new，用make
//2.map是引用类型
//3.map容量:和数组不同，map 可以根据新增的 key-value 对动态的伸缩，因此它不存在固定长度或者最大限制。
//4.但是你也可以选择标明 map 的初始容量 capacity，就像这样：make(map[keytype]valuetype，cap)。
func main() {
	test()
	//badCase()
	//case1()
	//case2()
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

func badCase() {
	//决不能用new创建map，得到的是一个空引用的指针，相当于声明了一个未初始化的变量并取了地址
	mapCreated := new(map[string]float32)
	log.Println(mapCreated, &mapCreated)
	//mapCreated["ss"] = 33.33
}

func test() {
	//intList := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	//sort.Ints(intList)
	//log.Println(intList)

	barVal := map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}

	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}
	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v / ", k, barVal[k])
	}

}
