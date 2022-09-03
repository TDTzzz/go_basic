package main

import "fmt"

//值接收者 指针接收者
//无论是类型的值或指针，在调用指针接收者方法或值接收者方法时，编译器都对他们做了相应的转换，所以互相混着调用都能成功。

type Duck struct {
	legs    int
	hasWing bool
	canSwim bool
}

// 值接收者方法
func (dk Duck) SetLegs(num int) {
	dk.legs = num
}

// 指针接收者方法
func (dk *Duck) ChangeLegs(num int) {
	dk.legs = num
	// (*dk).legs = num
}

func main() {
	first := Duck{}
	fmt.Println("1. First duck legs: ", first.legs)
	first.SetLegs(2)
	fmt.Println("2. First duck legs: ", first.legs)
	first.ChangeLegs(4)
	fmt.Println("3. First duck legs: ", first.legs)

	second := new(Duck)
	fmt.Println("4. Second duck legs: ", second.legs)
	second.SetLegs(4)
	fmt.Println("5. Second duck legs: ", second.legs)
	second.ChangeLegs(4)
	fmt.Println("6. Second duck legs: ", second.legs)
}
