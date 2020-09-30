package main

import "fmt"

func main() {
	//数组是值语义，一个数组变量即表示整个数组，并不是隐式的指向第一个元素的指针
	//数组指针类型除了类型和数组不同之外，通过数组指针操作数组的方式和通过数组本身的操作类似
	var a = [...]int{1, 2, 3}
	b := &a
	a[0] = 22
	fmt.Println(a[0], a[1])
	fmt.Println(b[0], b[1])

	//空数组在内存中不占用空间  可以用于 管道的同步操作

	var arr1 = new([5]int)
	arr := arr1
	arr1[2] = 100
	fmt.Println(arr1[2], arr[2])

	var arr2 [5]int
	newarr := arr2
	arr2[2] = 100
	fmt.Println(arr2[2], newarr[2])
}
