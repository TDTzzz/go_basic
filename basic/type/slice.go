package main

import "fmt"

//切片可以和nil进行比较，只有当切片底层数据指针为空时切片本身为nil，这时候切片的长度和容量信息将是无效的

func main() {
	//var (
	//	a []int               // nil切片, 和 nil 相等, 一般用来表示一个不存在的切片
	//	b = []int{}           // 空切片, 和 nil 不相等, 一般用来表示一个空的集合
	//	c = []int{1, 2, 3}    // 有3个元素的切片, len和cap都为3
	//	d = c[:2]             // 有2个元素的切片, len为2, cap为3
	//	e = c[0:2:cap(c)]     // 有2个元素的切片, len为2, cap为3
	//	f = c[:0]             // 有0个元素的切片, len为0, cap为3
	//	g = make([]int, 3)    // 有3个元素的切片, len和cap都为3
	//	h = make([]int, 2, 3) // 有2个元素的切片, len为2, cap为3
	//	i = make([]int, 0, 3) // 有0个元素的切片, len为0, cap为3
	//)

	var b = []int{0: 1}
	fmt.Println(cap(b))
	b = append(b, []int{1, 2, 3}...)
	fmt.Println(b)
	fmt.Println(cap(b))

	//也可以从slice的开头加元素
	//但是开头一般会导致内存的重新分配,性能会比从尾部差
	var a = []int{1, 2, 3}
	a = append([]int{0}, a...)
	fmt.Println(a)

	//DEL 删除切片

	a = []int{1, 2, 3}
	a = a[:len(a)-1] // 删除尾部1个元素
	//a = a[:len(a)-N]   // 删除尾部N个元素、

	a = []int{1, 2, 3}
	a = a[1:] // 删除开头1个元素
	//a = a[N:] // 删除开头N个元素

	//原地完成（所谓原地完成是指在原有的切片数据对应的内存区间内完成，不会导致内存空间结构的变化）
	a = []int{1, 2, 3}
	a = append(a[:0], a[1:]...) // 删除开头1个元素
	//a = append(a[:0], a[N:]...) // 删除开头N个元素

	//用copy完成
	a = []int{1, 2, 3}
	a = a[:copy(a, a[1:])] // 删除开头1个元素
	//a = a[:copy(a, a[N:])] // 删除开头N个元素

	aa := [5]int{1, 2, 3, 4, 5}
	tt := aa[1:3:5] //cap:5-1=4   len:2
	fmt.Println(len(tt), cap(tt))
}

//slice是引用类型，在不扩容的情况下，改变底层同数组的slice，同数组的其他slice也会相应的改变
//但是当扩容超出原有容量时，会复制数据到新数组。所以改数据，不会像之前一样影响到其他同数组slice
func goodCase() {
	s1 := []int{1, 2, 3}
	fmt.Println(len(s1), cap(s1), s1) // 输出 3 3 [1 2 3]
	s2 := s1[1:]
	fmt.Println(len(s2), cap(s2), s2) // 输出 2 2 [2 3]
	for i := range s2 {
		s2[i] += 20
	}
	// s2的修改会影响到数组数据，s1输出新数据
	fmt.Println(s1) // 输出 [1 22 23]
	fmt.Println(s2) // 输出 [22 23]

	s2 = append(s2, 4) // append  s2容量为2，这个操作导致了切片 s2扩容，会生成新的底层数组。

	for i := range s2 {
		s2[i] += 10
	}
	// s1 的数据现在是老数据，而s2扩容了，复制数据到了新数组，他们的底层数组已经不是同一个了。
	fmt.Println(len(s1), cap(s1), s1) // 输出3 3 [1 22 23]
	fmt.Println(len(s2), cap(s2), s2) // 输出3 4 [32 33 14]
}
