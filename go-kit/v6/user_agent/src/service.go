package src

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"go_basic/go-kit/v6/user_agent/pb"
	"go_basic/go-kit/v6/utils"
)

type Service interface {
	Login(ctx context.Context, in *pb.Login) (ack *pb.LoginAck, err error)
}

type baseServer struct {
	logger *zap.Logger
}

func (s baseServer) Login(ctx context.Context, in *pb.Login) (ack *pb.LoginAck, err error) {
	if in.Account != "tdtzzz" || in.Password != "123456" {
		err = errors.New("用户信息错误")
		return
	}
	ack = &pb.LoginAck{}
	ack.Token, err = utils.CreateJwtToken(in.Account, 1)
	return
}

func NewService(log *zap.Logger) Service {
	var server Service
	server = &baseServer{log}
	//middleware
	return server
}
