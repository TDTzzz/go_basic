package v2_endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
	"go_basic/go-kit/v2/v2_service"
)

type EndPointServer struct {
	AddEndpoint endpoint.Endpoint
}

func MakeEndpointServer(svc v2_service.Service, log *zap.Logger) EndPointServer {
	var addEndPoint endpoint.Endpoint
	{
		addEndPoint = MakeAddEndPoint(svc)
		addEndPoint = LoggingMiddleware(log)(addEndPoint)
	}
	return EndPointServer{AddEndpoint: addEndPoint}
}

func MakeAddEndPoint(svc v2_service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(v2_service.TestRequest)
		out := svc.Test(ctx, req)
		return out, nil
	}
}

func (s EndPointServer) Test(ctx context.Context, in v2_service.TestRequest) v2_service.TestResponse {
	res, _ := s.AddEndpoint(ctx, in)
	return res.(v2_service.TestResponse)
}
