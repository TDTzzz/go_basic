package main

import (
	"fmt"
	"go_basic/go_kit_project/demo/endpoint"
	"go_basic/go_kit_project/demo/service"
	"go_basic/go_kit_project/demo/transport"
	"net/http"
)

func main() {

	server := service.NewService()
	endpoints := endpoint.NewEndPointServer(server)
	httpHandler := transport.NewHttpHandler(endpoints)

	fmt.Println("server run localhost:8888")
	_ = http.ListenAndServe("127.0.0.1:8888", httpHandler)
}
