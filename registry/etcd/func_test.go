package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"testing"
	"time"
)

func TestPutGGet(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed,err%v", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	//put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "name", "tdtzzz")
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed,err %v\n", err)
		return
	}
	//get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "name")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed,err %v\n", err)
		return
	}
	t.Log("----！！！！！")
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
}
