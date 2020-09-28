package src

import (
	"context"
	"errors"
	"github.com/go-kit/kit/metrics"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"go_basic/go-kit/v9/user_agent/pb"
	"go_basic/go-kit/v9/utils"
)

func init() {

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
	ack = &pb.LoginAck{}
	ack.Token, err = utils.CreateJwtToken(in.Account, 1)
	return
}

func NewService(log *zap.Logger, counter metrics.Counter, histogram metrics.Histogram, tracer opentracing.Tracer) Service {
	var server Service
	//各种中间件
	return server
}
