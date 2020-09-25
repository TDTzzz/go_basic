package main

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"go_basic/registry/etcd/demo"
	"time"
)

func main() {
	options := demo.NewOptions("svc.test", 10, clientv3.Config{
		Endpoints:   []string{"http://localhost:2379/"},
		DialTimeout: 5 * time.Second,
	})

	for i := 0; i <= 3; i++ {
		r, err := demo.NewRegistry(options)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = r.RegistryNode(demo.PutNode{Addr: fmt.Sprintf("127.0.0.1:%d%d%d%d", i, i, i, i)})
		if err != nil {
			fmt.Println(err)
			return
		}

		if i == 3 {
			go func() {
				time.Sleep(time.Second * 20)
				r.UnRegistry()
			}()
		}
	}

	time.Sleep(time.Hour * 5)
}
