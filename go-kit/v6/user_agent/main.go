package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/etcdv3"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"go.uber.org/zap"
	"go_basic/go-kit/v6/user_agent/pb"
	"go_basic/go-kit/v6/user_agent/src"
	"go_basic/go-kit/v6/utils"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var grpcAddr = flag.String("g", "127.0.0.1:8881", "grpcAddr")

var quitChan = make(chan error, 1)

func main() {
	flag.Parse()

	var (
		etcdAddres = []string{"127.0.0.1:2379"}
		serName    = "svc.user.agent"
		grpcAddr   = *grpcAddr
		ttl        = 5 * time.Second
	)
	utils.NewLoggerServer()
	//初始化etcd客户端
	options := etcdv3.ClientOptions{
		DialTimeout:   ttl,
		DialKeepAlive: ttl,
	}

	etcdClient, err := etcdv3.NewClient(context.Background(), etcdAddres, options)

	if err != nil {
		utils.GetLogger().Error("[user_agent] NewClient", zap.Error(err))
		return
	}

	Register := etcdv3.NewRegistrar(etcdClient, etcdv3.Service{
		Key:   fmt.Sprint("%s/%s", serName, grpcAddr),
		Value: grpcAddr,
	}, log.NewNopLogger())

	go func() {
		golangLimit := rate.NewLimiter(10, 1)
		server := src.NewService(utils.GetLogger())
		endpoints := src.NewEndPointServer(server, golangLimit)
		grpcServer := src.NewGRPCServer(endpoints, utils.GetLogger())
		grpcListener, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			utils.GetLogger().Warn("[user_agent] Listen", zap.Error(err))
			quitChan <- err
			return
		}
		Register.Register()
		baseServer := grpc.NewServer(grpc.UnaryInterceptor(grpctransport.Interceptor))
		pb.RegisterUserServer(baseServer, grpcServer)
		if err = baseServer.Serve(grpcListener); err != nil {
			utils.GetLogger().Warn("[user_agent] Serve", zap.Error(err))
			quitChan <- err
			return
		}
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		quitChan <- fmt.Errorf("%s", <-c)
	}()
	utils.GetLogger().Info("[user_agent] run" + grpcAddr)
	err = <-quitChan
	Register.Deregister()
	utils.GetLogger().Info("[user_agent] quit err", zap.Error(err))
}
