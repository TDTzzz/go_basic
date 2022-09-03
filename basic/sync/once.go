package main

import "sync"

type singleton struct{}

var instance *singleton
var once sync.Once

//once单例
func main() {

}

func getInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
