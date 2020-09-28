package src

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go_basic/go-kit/v9/user_agent/pb"
)

type traceMiddlewareServer struct {
	next   Service
	tracer opentracing.Tracer
}

func (l traceMiddlewareServer) Login(ctx context.Context, in *pb.Login) (out *pb.LoginAck, err error) {
	span, ctxContent := opentracing.StartSpanFromContextWithTracer(ctx, l.tracer, "service", opentracing.Tag{
		Key:   string(ext.Component),
		Value: "NewTraceServerMiddleware",
	})
	defer func() {
		span.LogKV("account", in.GetAccount(), "password", in.GetPassword())
		span.Finish()
	}()
	out, err = l.next.Login(ctxContent, in)
	return
}

func NewTraceMiddlewareServer(tracer opentracing.Tracer) NewMiddlewareServer {
	return func(service Service) Service {
		return traceMiddlewareServer{
			next:   service,
			tracer: tracer,
		}
	}
}
