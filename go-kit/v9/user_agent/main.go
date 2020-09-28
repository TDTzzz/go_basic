package user_agent

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	metricsprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/kit/sd/etcdv3"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"go_basic/go-kit/v9/user_agent/pb"
	"go_basic/go-kit/v9/user_agent/src"
	"go_basic/go-kit/v9/utils"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"hash/crc32"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
		return
	}
	Registrar := etcdv3.NewRegistrar(etcdClient, etcdv3.Service{
		Key:   fmt.Sprintf("%s/%d", serName, crc32.ChecksumIEEE([]byte(*grpcAddr))),
		Value: *grpcAddr,
	}, log.NewNopLogger())

	go func() {
		tracer, _, err := utils.NewJaegerTracer("user_agent_server")
		if err != nil {
			utils.GetLogger().Warn("[user_agent] NewJaegerTracer",
				zap.Error(err))
			quitChan <- err
		}
		count := metricsprometheus.NewCounterFrom(prometheus.CounterOpts{
			Subsystem: "user_agent",
			Name:      "request_count",
			Help:      "Number of requests",
		}, []string{"method"})

		histogram := metricsprometheus.NewHistogramFrom(prometheus.HistogramOpts{
			Subsystem: "user_agent",
			Name:      "request_consume",
			Help:      "Request consumes time",
		}, []string{"method"})
		golangLimit := rate.NewLimiter(10, 1)
		server := src.NewService(utils.GetLogger(), count, histogram, tracer)
		endpoints := src.NewEndPointServer(server, golangLimit, tracer)
		grpcServer := src.NewGRPCServer(endpoints, utils.GetLogger())
		grpcListener, err := net.Listen("tcp", *grpcAddr)
		if err != nil {
			utils.GetLogger().Warn("[user_agent]grpc run" + *grpcAddr)
			quitChan <- err
			return
		}
		Registrar.Register()
		utils.GetLogger().Info("[user_agent] grpc run" + *grpcAddr)
		chainUnaryServer := grpcmiddleware.ChainUnaryServer(
			grpctransport.Interceptor,
			grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(tracer)),
			grpc_zap.UnaryServerInterceptor(utils.GetLogger()),
		)
		baseServer := grpc.NewServer(grpc.UnaryInterceptor(chainUnaryServer))
		pb.RegisterUserServer(baseServer, grpcServer)
		quitChan <- baseServer.Serve(grpcListener)
	}()
	go func() {
		utils.GetLogger().Info("[user_agent] prometheus run" + *prometheusAddr)
		m := http.NewServeMux()
		m.Handle("/metrics", promhttp.Handler())
		quitChan <- http.ListenAndServe(*prometheusAddr, m)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		quitChan <- fmt.Errorf("%s", <-c)
	}()

	err = <-quitChan
	Registrar.Deregister()
	utils.GetLogger().Info("[user_agent] quit", zap.Any("info", err))
}
