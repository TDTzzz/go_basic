package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go_basic/go_kit_project/demo/service"
)

//EndPoint方法集合
type EndPointServer struct {
	DemoEndPoint endpoint.Endpoint
}

func NewEndPointServer(svc service.DemoService) EndPointServer {
	var demoEndPoint endpoint.Endpoint
	{
		demoEndPoint = MakeDemoEndPoint(svc)
	}
	return EndPointServer{DemoEndPoint: demoEndPoint}
}

func (s EndPointServer) HealthCheck(req service.DemoReq) service.DemoRes {
	res := s.HealthCheck(req)
	return res
}

//把service的方法封装到Endpoint中
func MakeDemoEndPoint(s service.DemoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(service.DemoReq)
		res := s.HealthCheck(req)
		return res, nil
	}
}
