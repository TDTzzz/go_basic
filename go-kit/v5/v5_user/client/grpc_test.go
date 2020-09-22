package client

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"go_basic/go-kit/v5/v5_user/pb"
	"go_basic/go-kit/v5/v5_user/v5_service"
	"go_basic/pkg_library/logtool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"testing"
)

//go-kit客户端调用方法
func TestGrpcClient(t *testing.T) {
	logger := logtool.NewLogger(
		logtool.SetAppName("go-kit"),
		logtool.SetDevelopment(true),
		logtool.SetLevel(zap.DebugLevel),
	)
	conn, err := grpc.Dial("127.0.0.1:8881", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.Close()
	svr := NewGRPCClient(conn, logger)
	ack, err := svr.Login(context.Background(), &pb.Login{
		Account:  "tdtzzz",
		Password: "123456",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ack.Token)
}

//grpc原生客户端
func TestGrpc(t *testing.T) {
	serviceAddress := "127.0.0.1:8881"
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		panic("connect error")
	}
	defer conn.Close()
	userClient := pb.NewUserClient(conn)
	UUID := uuid.NewV5(uuid.Must(uuid.NewV4(), nil), "req_uuid").String()
	md := metadata.Pairs(v5_service.ContextReqUUid, UUID)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	res, err := userClient.RpcUserLogin(ctx, &pb.Login{
		Account:  "tdtzzz",
		Password: "123456",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(res.Token)
}
