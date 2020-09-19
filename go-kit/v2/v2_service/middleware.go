package v2_service

import (
	"context"
	"fmt"
	"go.uber.org/zap"
)

const ContextReqUUid = "req_uuid"

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

func (l logMiddlewareServer) Test(ctx context.Context, in TestRequest) (out TestResponse) {
	defer func() {
		l.log.Debug(fmt.Sprint(ctx.Value(ContextReqUUid)),
			zap.Any("调用v2_service logMiddlewareServer", "Test"),
			zap.Any("req", in),
			zap.Any("res", out),
		)
	}()
	out = l.next.Test(ctx, in)
	return out
}
