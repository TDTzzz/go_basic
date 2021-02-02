package main

import "fmt"

//pipeLine
func main() {
	coms := buy(10)
	phones := build(coms)
	packs := pack(phones)

	for p := range packs {
		fmt.Println(p)
	}
}

//工序一：买东西
func buy(n int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 1; i <= n; i++ {
			out <- fmt.Sprint("配件", i)
		}
	}()
	return out
}

//工序二：组装
func build(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "组装（" + c + ")"
		}
	}()
	return out
}

//工序三：打包
func pack(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "打包（" + c + "）"
		}
	}()
	return out
}
