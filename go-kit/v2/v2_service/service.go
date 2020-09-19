package v2_service

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"time"
)

type Service interface {
	Test(ctx context.Context, in TestRequest) TestResponse
}

type baseService struct {
	log *zap.Logger
}

func (s baseService) Test(ctx context.Context, in TestRequest) TestResponse {
	//模拟耗时
	time.Sleep(time.Millisecond * 2)
	s.log.Debug(fmt.Sprint(ctx.Value(ContextReqUUid),
		zap.Any("调用v2_service Service", "Test 处理请求")))
	out := in.A + in.B
	s.log.Debug(fmt.Sprint(ctx.Value(ContextReqUUid)),
		zap.Any("调用v2_service Service", "Test 处理请求"),
		zap.Any("处理返回值", out))
	return TestResponse{Res: out}
}

func NewServer(log *zap.Logger) Service {
	var server Service
	server = &baseService{log: log}
	server = NewLogMiddlewareServer(log)(server)
	return server
}
