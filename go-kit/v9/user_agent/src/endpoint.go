package src

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/opentracing/opentracing-go"
	"go_basic/go-kit/v9/user_agent/pb"
	"golang.org/x/time/rate"
)

type EndPointServer struct {
	LoginEndPoint endpoint.Endpoint
}

func (s EndPointServer) Login(ctx context.Context, in *pb.Login) (ack *pb.LoginAck, err error) {
	res, err := s.LoginEndPoint(ctx, in)
	if err != nil {
		fmt.Println("s.LoginEndPoint", err)
		return nil, err
	}
	return res.(*pb.LoginAck), nil
}

func NewEndPointServer(svc Service, limit *rate.Limiter, tracer opentracing.Tracer) EndPointServer {
	var loginEndPoint endpoint.Endpoint
	{
		loginEndPoint = MakeLoginEndPoint(svc)

	}

	return EndPointServer{LoginEndPoint: loginEndPoint}
}

func MakeLoginEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.Login)
		return s.Login(ctx, req)
	}
}
