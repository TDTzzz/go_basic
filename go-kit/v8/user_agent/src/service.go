package src

import (
	"context"
	"errors"
	"github.com/go-kit/kit/metrics"
	"go.uber.org/zap"
	"go_basic/go-kit/v8/user_agent/pb"
	"go_basic/go-kit/v8/utils"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Service interface {
	Login(ctx context.Context, in *pb.Login) (ack *pb.LoginAck, err error)
}

type baseServer struct {
	logger *zap.Logger
}

func (b baseServer) Login(ctx context.Context, in *pb.Login) (ack *pb.LoginAck, err error) {
	if in.Account != "tdtzzz" || in.Password != "123456" {
		err = errors.New("用户信息错误")
		return
	}
	//模拟耗时
	//模拟错误

	//if rand.Intn(10) > 3 {
	//	err = errors.New("服务器运行错误")
	//	return
	//}
	ack = &pb.LoginAck{}
	ack.Token, err = utils.CreateJwtToken(in.Account, 1)
	return
}

func NewService(log *zap.Logger, counter metrics.Counter, histogram metrics.Histogram) Service {
	var server Service
	server = &baseServer{log}
	server = NewLogMiddlewareServer(log)(server)
	server = NewMetricsMiddlewareServer(counter, histogram)(server)
	return server
}
