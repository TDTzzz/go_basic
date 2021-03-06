package v5_service

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"go_basic/go-kit/v5/utils"
	"go_basic/go-kit/v5/v5_user/pb"
)

type Service interface {
	Login(ctx context.Context, in *pb.Login) (ack *pb.LoginAck, err error)
}

type baseServer struct {
	logger *zap.Logger
}

func (s baseServer) Login(ctx context.Context, in *pb.Login) (ack *pb.LoginAck, err error) {
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUid)),
		zap.Any("调用v5_service Service", "Login 处理请求"))
	if in.Account != "tdtzzz" || in.Password != "123456" {
		err = errors.New("用户信息错误")
		return
	}
	ack = &pb.LoginAck{}
	ack.Token, err = utils.CreateJwtToken(in.Account, 1)
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUid)),
		zap.Any("调用v5_service Service", "Login 处理请求"),
		zap.Any("处理返回值", ack))
	return
}

func NewService(log *zap.Logger) Service {
	var server Service
	server = &baseServer{log}
	server = NewLogMiddlewareServer(log)(server)
	return server
}
