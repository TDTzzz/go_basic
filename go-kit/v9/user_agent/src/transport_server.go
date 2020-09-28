package src

import (
	"context"
	"fmt"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"go.uber.org/zap"
	"go_basic/go-kit/v9/user_agent/pb"
	"go_basic/go-kit/v9/utils"
	"google.golang.org/grpc/metadata"
)

type grpcServcer struct {
	login grpctransport.Handler
}

func NewGRPCServer(endpoint EndPointServer, log *zap.Logger) pb.UserServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
			ctx = context.WithValue(ctx, utils.ContextReqUUid, md.Get(utils.ContextReqUUid))
			return ctx
		}),
		grpctransport.ServerErrorHandler(utils.NewZapLogErrorHandler(log)),
	}
	return &grpcServcer{login: grpctransport.NewServer(
		endpoint.LoginEndPoint,
		RequestGrpcLogin,
		ResponseGrpcLogin,
		options...
	)}
}

func (s *grpcServcer) RpcUserLogin(ctx context.Context, req *pb.Login) (*pb.LoginAck, error) {
	_, rep, err := s.login.ServeGRPC(ctx, req)
	if err != nil {
		fmt.Println("s.login.ServeGRPC", err)
		return nil, err
	}
	return rep.(*pb.LoginAck), nil
}

func RequestGrpcLogin(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.Login)
	return &pb.Login{
		Account:  req.GetAccount(),
		Password: req.GetPassword(),
	}, nil
}

func ResponseGrpcLogin(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.LoginAck)
	return &pb.LoginAck{
		Token: resp.Token,
	}, nil
}
