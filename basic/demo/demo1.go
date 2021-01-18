package main

import (
	"fmt"
	"sync"
)

func main() {
	case3()
}

//结果不是恒定的，count++非原子性
func case1() {
	count := 0
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				count++
			}
		}()
	}
	wg.Wait()

	fmt.Println(count)
}

//也是非固定的，因为go虽然有序创建了，但goroutine不一定开始运行了
// 第一次输出
//0
//1
//2
//4
//3
//
//// 第二次输出
//4
//0
//1
//2
//3
func case2() {
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}

// 第一次输出
//5
//5
//5
//5
//5
// 多输出几次
//5
//3
//5
//5
//5
func case3() {
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()
}
