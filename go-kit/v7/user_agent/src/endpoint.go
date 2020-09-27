package src

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"go_basic/go-kit/v7/user_agent/pb"
	"golang.org/x/time/rate"
)

type EndPointServer struct {
	LoginEnPoint endpoint.Endpoint
}

func NewEndPointServer(svc Service, limit *rate.Limiter) EndPointServer {
	var loginEndPoint endpoint.Endpoint
	{
		loginEndPoint = MakeLoginEndPoint(svc)
		loginEndPoint = NewGolangRateAllowMiddleware(limit)(loginEndPoint)
	}
	return EndPointServer{LoginEnPoint: loginEndPoint}
}

func (s EndPointServer) Login(ctx context.Context, in *pb.Login) (*pb.LoginAck, error) {
	res, err := s.LoginEnPoint(ctx, in)
	if err != nil {
		fmt.Println("s.LoginEndPoint", err)
		return nil, err
	}
	return res.(*pb.LoginAck), nil
}

func MakeLoginEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.Login)
		return s.Login(ctx, req)
	}
}
