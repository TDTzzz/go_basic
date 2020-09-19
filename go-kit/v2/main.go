package main

import (
	"go_basic/go-kit/v2/utils"
	"go_basic/go-kit/v2/v2_endpoint"
	"go_basic/go-kit/v2/v2_service"
	"go_basic/go-kit/v2/v2_transport"
	"net/http"
)

func main() {
	utils.NewLoggerServer()
	server := v2_service.NewServer(utils.GetLogger())
	endpoints := v2_endpoint.MakeEndpointServer(server, utils.GetLogger())
	httpHandler := v2_transport.NewHttpHandler(endpoints, utils.GetLogger())
	http.ListenAndServe("0.0.0.0:8886", httpHandler)
}
