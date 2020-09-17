package main

import (
	"context"
	"fmt"
	endpoint2 "go_basic/go-kit/v1/endpoint"
	"go_basic/go-kit/v1/service"
	"go_basic/go-kit/v1/transport"
	"net/http"
)

func main() {
	ctx := context.Background()

	var svc service.Service
	svc = service.ArithmeticService{}

	endpoint := endpoint2.MakeArithmeticEndpoint(svc)

	r := transport.MakeHttpHandler(ctx, endpoint)
	fmt.Println("server run 0.0.0.0:8886")
	http.ListenAndServe("0.0.0.0:8886", r)
}
