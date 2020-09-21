package v4_endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	"go_basic/go-kit/v4/v4_service"
	"golang.org/x/time/rate"
)

type EndPointServer struct {
	AddEndPoint   endpoint.Endpoint
	LoginEndPoint endpoint.Endpoint
}

func NewEndPointServer(svc v4_service.Service, log *zap.Logger, limit *rate.Limiter, limiter ratelimit.Limiter) EndPointServer {
	var addEndPoint endpoint.Endpoint
	{
		addEndPoint = MakeAddEndPoint(svc)
		addEndPoint = LoggingMiddleware(log)(addEndPoint)
		addEndPoint = AuthMiddleware(log)(addEndPoint)
		addEndPoint = NewUberRateMiddleware(limiter)(addEndPoint)
	}

	var loginEndPoint endpoint.Endpoint
	{
		loginEndPoint = MakeLoginEndPoint(svc)
		loginEndPoint = LoggingMiddleware(log)(loginEndPoint)
		loginEndPoint = NewGolangRateAllowMiddleware(limit)(loginEndPoint)
	}
	return EndPointServer{
		AddEndPoint:   addEndPoint,
		LoginEndPoint: loginEndPoint,
	}
}

func MakeAddEndPoint(s v4_service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(v4_service.Add)
		res := s.TestAdd(ctx, req)
		return res, nil
	}
}

func MakeLoginEndPoint(s v4_service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(v4_service.Login)
		return s.Login(ctx, req)
	}
}

func (s EndPointServer) TestAdd(ctx context.Context, in v4_service.Add) v4_service.AddAck {
	res, _ := s.AddEndPoint(ctx, in)
	return res.(v4_service.AddAck)
}

func (s EndPointServer) Login(ctx context.Context, in v4_service.Login) (ack v4_service.LoginAck, err error) {
	res, err := s.LoginEndPoint(ctx, in)
	return res.(v4_service.LoginAck), err
}
