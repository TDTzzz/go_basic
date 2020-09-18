package v2_service

import (
	"go.uber.org/zap"
)

type NewMiddlewareServer func(Service) Service

type logMiddlewareServer struct {
	log  *zap.Logger
	next Service
}

func NewLogMiddlewareServer(log *zap.Logger) NewMiddlewareServer {
	return func(service Service) Service {
		return logMiddlewareServer{
			log:  log,
			next: service,
		}
	}
}

func (l logMiddlewareServer) Test(in TestRequest) (out TestResponse) {
	defer func() {
		l.log.Debug("kit_log_test_v2",
			zap.Any("调用v2_service logMiddlewareServer", "Add"),
			zap.Any("req", in),
			zap.Any("res", out),
		)
	}()
	out = l.next.Test(in)
	return out
}
