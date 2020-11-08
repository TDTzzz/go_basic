package transport

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"go_basic/go_kit_project/demo/endpoint"
	"go_basic/go_kit_project/demo/service"
	"net/http"
	"strconv"
)

func NewHttpHandler(endpoint endpoint.EndPointServer) http.Handler {
	m := http.NewServeMux()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(httptransport.DefaultErrorEncoder),
	}

	m.Handle("/health", httptransport.NewServer(
		endpoint.DemoEndPoint,
		decodeHttpDemoReq,
		encodeHttpDemoRes,
		options...,
	))

	return m
}

//把请求转换成Server所需的格式
func decodeHttpDemoReq(_ context.Context, r *http.Request) (interface{}, error) {
	var (
		req service.DemoReq
		err error
	)

	req.Port, err = strconv.Atoi(r.FormValue("port"))
	if err != nil {
		return req, err
	}
	req.Ip = r.FormValue("ip")
	return req, nil
}

//把返回值返回给用户
func encodeHttpDemoRes(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
