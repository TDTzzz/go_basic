package main

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"go.uber.org/zap"
	"go_basic/go-kit/v5/utils"
	"go_basic/go-kit/v5/v5_user/pb"
	"go_basic/go-kit/v5/v5_user/v5_endpoint"
	"go_basic/go-kit/v5/v5_user/v5_service"
	"go_basic/go-kit/v5/v5_user/v5_transport"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	utils.NewLoggerServer()
	golangLimit := rate.NewLimiter(10, 1)
	server := v5_service.NewService(utils.GetLogger())
	endpoints := v5_endpoint.NewEndPointServer(server, utils.GetLogger(), golangLimit)
	grpcServer := v5_transport.NewGRPCServer(endpoints, utils.GetLogger())
	utils.GetLogger().Info("server run :8881")
	grpcListener, err := net.Listen("tcp", ":8881")
	if err != nil {
		utils.GetLogger().Warn("Listen", zap.Error(err))
		os.Exit(0)
	}

	baseServer := grpc.NewServer(grpc.UnaryInterceptor(grpctransport.Interceptor))
	pb.RegisterUserServer(baseServer, grpcServer)
	if err = baseServer.Serve(grpcListener); err != nil {
		utils.GetLogger().Warn("Serve", zap.Error(err))
	}
}
