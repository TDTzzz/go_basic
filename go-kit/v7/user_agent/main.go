package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/etcdv3"
	"go.uber.org/zap"
	"go_basic/go-kit/v7/utils"
	"hash/crc32"
	"time"
)

var grpcAddr = flag.String("g", "127.0.0.1:8881", "grpcAddr")
var prometheusAddr = flag.String("p", "192.168.2.28:10001", "prometheus addr")

var quitChan = make(chan error, 1)

func main() {
	flag.Parse()

	var (
		etcdAddrs = []string{"127.0.0.1:2379"}
		serName   = "svc.user.agent"
		ttl       = 5 * time.Second
	)
	utils.NewLoggerServer()
	//初始化etcd客户端
	options := etcdv3.ClientOptions{
		DialTimeout:   ttl,
		DialKeepAlive: ttl,
	}
	etcdClient, err := etcdv3.NewClient(context.Background(), etcdAddrs, options)
	if err != nil {
		utils.GetLogger().Error("[user_agent] NewClient", zap.Error(err))
	}
	Registrar := etcdv3.NewRegistrar(etcdClient, etcdv3.Service{
		Key:   fmt.Sprintf("%s/%d", serName, crc32.ChecksumIEEE([]byte(*grpcAddr))),
		Value: *grpcAddr,
	}, log.NewNopLogger())

	go func() {
	}()
}
