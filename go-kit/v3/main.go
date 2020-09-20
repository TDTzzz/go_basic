package main

import (
	"go_basic/go-kit/v3/utils"
	"go_basic/go-kit/v3/v3_endpoint"
	"go_basic/go-kit/v3/v3_service"
	"go_basic/go-kit/v3/v3_transport"
	"net/http"
)

func main() {
	utils.NewLoggerServer()
	server := v3_service.NewService(utils.GetLogger())
	endpoints := v3_endpoint.NewEndPointServer(server, utils.GetLogger())
	httpHandler := v3_transport.NewHttpHandler(endpoints, utils.GetLogger())
	utils.GetLogger().Info("server run 0.0.0.0:8887")
	_ = http.ListenAndServe(":8887", httpHandler)
}
