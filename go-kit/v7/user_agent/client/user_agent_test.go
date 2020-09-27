package client

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	uuid "github.com/satori/go.uuid"
	"go_basic/go-kit/v7/user_agent/pb"
	"go_basic/go-kit/v7/user_agent/src"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"os"
	"testing"
)

func TestNewUserAgentClient(t *testing.T) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	client, err := NewUserAgentClient([]string{"127.0.0.1:2379"}, logger)
	if err != nil {
		t.Error(err)
		return
	}
	for i := 0; i < 20; i++ {
		userAgent, err := client.UserAgentClient()
		if err != nil {
			t.Error(err)
			return
		}
		ack, err := userAgent.Login(context.Background(), &pb.Login{
			Account:  "tdtzzz",
			Password: "123456",
		})
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(ack.Token)
	}

}

func TestGrpc(t *testing.T) {
	serviceAddress := "127.0.0.1:8881"
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		panic("connect error")
	}
	defer conn.Close()
	userClient := pb.NewUserClient(conn)
	UUID := uuid.NewV5(uuid.Must(uuid.NewV4(), nil), "req_uuid").String()
	md := metadata.Pairs(src.ContextReqUUid, UUID)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	for i := 0; i < 20; i++ {
		res, err := userClient.RpcUserLogin(ctx, &pb.Login{
			Account:  "hwholiday",
			Password: "123456",
		})
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println(res.Token)
	}
}
