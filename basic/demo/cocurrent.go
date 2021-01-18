//package main
//
//import (
//	"fmt"
//	"sync"
//)
//
//var mu sync.Mutex
//var chain string
//
////A: 不能编译
////B: 输出 main --> A --> B --> C
////C: 输出 main
////D: panic
//func main() {
//	chain = "main"
//	//A()
//	B()
//	fmt.Println(chain)
//}
//func A() {
//	mu.Lock()
//	defer mu.Unlock()
//	chain = chain + " --> A"
//	B()
//}
//func B() {
//	chain = chain + " --> B"
//	C()
//}
//func C() {
//	mu.Lock()
//	defer mu.Unlock()
//	chain = chain + " --> C"
//}

package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.RWMutex
var count int

func main() {
	go A()
	time.Sleep(2 * time.Second)
	mu.Lock()
	defer mu.Unlock()
	count++
	fmt.Println(count)
}
func A() {
	mu.RLock()
	defer mu.RUnlock()
	B()
}
func B() {
	time.Sleep(5 * time.Second)
	C()
}
func C() {
	mu.RLock()
	defer mu.RUnlock()
}
