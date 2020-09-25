package demo

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"testing"
	"time"
)

func TestNewRegistry(t *testing.T) {
	var op = Options{
		name: "svc.info",
		ttl:  10,
		config: clientv3.Config{
			Endpoints:   []string{"http://localhost:2379/"},
			DialTimeout: 5 * time.Second},
	}
	for i := 1; i <= 3; i++ {
		t.Log("!!!!!!!", i)
		r, err := NewRegistry(op)
		if err != nil {
			t.Error(err)
			return
		}
		err = r.RegistryNode(PutNode{Addr: fmt.Sprintf("127.0.0.1:%d%d%d%d", i, i, i, i)})
		if err != nil {
			t.Error(err)
			return
		}
		if i == 3 {
			go func() {
				time.Sleep(time.Second * 20)
				r.UnRegistry()
			}()
		}

	}
	time.Sleep(time.Second * 5)
}
