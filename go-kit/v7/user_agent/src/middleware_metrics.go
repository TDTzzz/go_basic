package src

import (
	"context"
	"github.com/go-kit/kit/metrics"
	"go_basic/go-kit/v7/user_agent/pb"
	"time"
)

type metricsMiddlewareServer struct {
	next      Service
	counter   metrics.Counter
	histogram metrics.Histogram
}

func (l metricsMiddlewareServer) Login(ctx context.Context, in *pb.Login) (out *pb.LoginAck, err error) {
	defer func(start time.Time) {
		method := []string{"method", "login"}
		l.counter.With(method...).Add(1)
		l.histogram.With(method...).Observe(time.Since(start).Seconds())
	}(time.Now())
	out, err = l.next.Login(ctx, in)
	return
}

func NewMetricsMiddlewareServer(counter metrics.Counter, histogram metrics.Histogram) NewMiddlewareServer {
	return func(service Service) Service {
		return metricsMiddlewareServer{
			next:      service,
			counter:   counter,
			histogram: histogram,
		}
	}
}
