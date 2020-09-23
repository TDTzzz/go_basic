package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/sd/etcdv3"
	"go_basic/go-kit/v6/utils"
	"time"
)

var grpcAddr = flag.String("g", "127.0.0.1:8881", "grpcAddr")

func main() {
	flag.Parse()
	var (
		etcdAddres = []string{"127.0.0.1:2379"}
		//serName    = "svc.user.agent"
		//grpcAddr   = *grpcAddr
		ttl = 5 * time.Second
	)

	utils.NewLoggerServer()
	//初始化etcd客户端
	options := etcdv3.ClientOptions{
		DialTimeout:   ttl,
		DialKeepAlive: ttl,
	}
	etcdClient, err := etcdv3.NewClient(context.Background(), etcdAddres, options)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(etcdClient)

}
