package main

import (
	"fmt"
	"sync"
	"time"
)

/**
互斥锁
使用Lock()加锁后，不能再继续对其加锁（同一个goroutine中，即：同步调用），否则会panic。只有在unlock()之后才能再次Lock()。
异步调用Lock()，是正当的锁竞争，当然不会有panic了。
适用于读写不确定场景，即读写次数没有明显的区别，并且只允许只有一个读或者写的场景，所以该锁也叫做全局锁。

func (m *Mutex) Unlock()用于解锁m，如果在使用Unlock()前未加锁，就会引起一个运行错误。
已经锁定的Mutex并不与特定的goroutine相关联，这样可以利用一个goroutine对其加锁，再利用其他goroutine对其解锁。
*/
func main() {
	//example_1()
	example_2()
}

func example_1() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(3)
	fmt.Println("Locking！！")
	mutex.Lock()
	fmt.Println("Locked！！")
	for i := 1; i < 4; i++ {
		go func(i int) {
			fmt.Println("Locking:", i)
			mutex.Lock()
			fmt.Println("Locked:", i)
			time.Sleep(3 * time.Second)
			mutex.Unlock()
			fmt.Println("unLock:", i)
			wg.Done()
		}(i)
	}
	time.Sleep(5 * time.Second)
	fmt.Println("ready unLock！！")
	mutex.Unlock()
	fmt.Println("unLock！！")
	wg.Wait()
}

//
type Dog struct {
	Name string
	L    *sync.Mutex
}

func (d *Dog) SetName(wg *sync.WaitGroup, name string) {
	defer func() {
		fmt.Println("unlock set name", name)
		d.L.Unlock()
		wg.Done()
	}()
	d.L.Lock()
	fmt.Println("lock set name:", name)
	time.Sleep(1 * time.Second)
	d.Name = name

}

func example_2() {
	wg := new(sync.WaitGroup)
	dog := Dog{}
	dog.L = &sync.Mutex{}
	names := []string{"aaa", "bbb", "ccc"}

	for _, name := range names {
		wg.Add(1)
		go dog.SetName(wg, name)
	}
	wg.Wait()
}
