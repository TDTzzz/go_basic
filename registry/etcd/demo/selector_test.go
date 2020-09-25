package demo

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"testing"
	"time"
)

func TestNewSelector(t *testing.T) {
	var op = SelectorOptions{
		name: "svc.info",
		config: clientv3.Config{
			Endpoints:   []string{"http://localhost:2379/"},
			DialTimeout: 5 * time.Second},
	}
	s, err := NewSelector(op)
	if err != nil {
		t.Error(err)
		return
	}
	for {
		val, err := s.Next()
		if err != nil {
			t.Error(err)
			continue
		}
		fmt.Println(val)
		time.Sleep(time.Second * 2)
	}
}
