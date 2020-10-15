package main

import (
	"fmt"
)

func main() {
	test_recover()

	fmt.Println("after recover")
}

func color() {
	for b := 40; b <= 47; b++ { // 背景色彩 = 40-47
		for f := 30; f <= 37; f++ { // 前景色彩 = 30-37
			for d := range []int{0, 1, 4, 5, 7, 8} { // 显示方式 = 0,1,4,5,7,8
				fmt.Printf(" %c[%d;%d;%dm%s(f=%d,b=%d,d=%d)%c[0m ", 0x1B, d, b, f, "", f, b, d, 0x1B)
			}
			fmt.Println("")
		}
		fmt.Println("")
	}
}

func test_recover() {
	defer func() {
		fmt.Println("defer func")
		if err := recover(); err != nil {
			fmt.Println("recover success")
		}
	}()

	arr := []int{1, 2, 3}
	fmt.Println(arr[4])
	fmt.Println("after panic")
}
