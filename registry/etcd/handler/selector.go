package main

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"go_basic/registry/etcd/demo"
	"log"
	"time"
)

func main() {
	selectorOp := demo.NewSelectorOptions("svc.test", clientv3.Config{
		Endpoints:   []string{"http://localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	s, err := demo.NewSelector(selectorOp)
	if err != nil {
		return
	}
	for {
		//fmt.Println("--")
		val, err := s.Next()
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(val)
		time.Sleep(time.Second * 2)
	}
}
