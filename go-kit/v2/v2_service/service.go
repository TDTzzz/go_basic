package v2_service

import "go.uber.org/zap"

type Service interface {
	Test(in TestRequest) TestResponse
}

type baseService struct {
	log *zap.Logger
}

func (s baseService) Test(in TestRequest) TestResponse {
	out := in.A + in.B
	return TestResponse{Res: out}
}

func NewServer(log *zap.Logger) Service {
	var server Service
	server = &baseService{log: log}
	server = NewLogMiddlewareServer(log)(server)
	return server
}
