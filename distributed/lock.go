package main

import (
	"log"
	"sync"
)

/**
分布式锁
https://github.com/chai2010/advanced-go-programming-book/blob/master/ch6-cloud/ch6-02-lock.md
 */

var noLockCnt int
var lockCnt int
var tryLockCnt int

//分布式锁
func main() {
	//noLockCase()
	//lockCase()
	tryLockCase()
}

//问题case
func noLockCase() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			noLockCnt++
		}()
	}
	wg.Wait()
	log.Println(noLockCnt)
}

//
func lockCase() {
	var wg sync.WaitGroup
	var lock sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lock.Lock()
			lockCnt++
			lock.Unlock()
		}()
	}
	wg.Wait()
	log.Println(lockCnt)
}

//tryLock:这种lock不像计数器那种，需要所有的goroutine都成功。
//而需要goroutine在抢锁失败后，放弃其流程

type Lock struct {
	c chan struct{}
}

func NewLock() Lock {
	var l Lock
	l.c = make(chan struct{}, 1)
	l.c <- struct{}{}
	return l
}

func (l Lock) Lock() bool {
	lockRes := false
	select {
	case <-l.c:
		lockRes = true
	default:
	}
	return lockRes
}

func (l Lock) UnLock() {
	l.c <- struct{}{}
}

func tryLockCase() {
	var wg sync.WaitGroup
	lock := NewLock()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !lock.Lock() {
				log.Println("lock失败")
				return
			}
			tryLockCnt++
			log.Println("当前的cnt", tryLockCnt)
			lock.UnLock()
		}()
		//time.Sleep(time.Second)
	}
	wg.Wait()
	log.Println("结果cnt", tryLockCnt)
}


/**
其他思路
1.远程调用redis，用setnx控制"抢占唯一资源"，效果和trylock类似
2.基于ZooKeeper
3.基于etcd
 */