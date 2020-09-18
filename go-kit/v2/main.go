package main

import (
	"go.uber.org/zap"
	"go_basic/go-kit/v2/v2_endpoint"
	"go_basic/go-kit/v2/v2_service"
	"go_basic/go-kit/v2/v2_transport"
	"net/http"
)

func main() {

	var log *zap.Logger
	server := v2_service.NewServer(log)
	endpoints := v2_endpoint.MakeEndpointServer(server, log)
	httpHandler := v2_transport.NewHttpHandler(endpoints, log)
	http.ListenAndServe("0.0.0.0:8885", httpHandler)
}
