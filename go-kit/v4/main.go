package main

import (
	"go.uber.org/ratelimit"
	"go_basic/go-kit/v4/utils"
	"go_basic/go-kit/v4/v4_endpoint"
	"go_basic/go-kit/v4/v4_service"
	"go_basic/go-kit/v4/v4_transport"
	"golang.org/x/time/rate"
	"net/http"
)

func main() {
	utils.NewLoggerServer()

	golangLimit := rate.NewLimiter(10, 1) //每秒产生10个令牌
	uberLimit := ratelimit.New(1)         //1秒请求一次
	server := v4_service.NewService(utils.GetLogger())
	endpoints := v4_endpoint.NewEndPointServer(server, utils.GetLogger(), golangLimit, uberLimit)
	transport := v4_transport.NewHttpHandler(endpoints, utils.GetLogger())

	utils.GetLogger().Info("server run :8887")
	_ = http.ListenAndServe(":8887", transport)
}
