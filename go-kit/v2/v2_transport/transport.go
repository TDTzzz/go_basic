package v2_transport

import (
	"context"
	"encoding/json"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	"go.uber.org/zap"
	"go_basic/go-kit/v2/v2_endpoint"
	"go_basic/go-kit/v2/v2_service"
	"log"
	"net/http"
	"strconv"
)

func NewHttpHandler(endpoint v2_endpoint.EndPointServer, log *zap.Logger) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(func(ctx context.Context, err error, w http.ResponseWriter) {
			log.Warn(fmt.Sprint("uuid"), zap.Error(err))
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(errWrapper{Error: err.Error()})
		}),
		httptransport.ServerBefore(func(ctx context.Context, request *http.Request) context.Context {
			log.Debug("添加uuid", zap.Any("UUID", "test_uuid"))
			ctx = context.WithValue(ctx, "qaa", "sss")
			return ctx
		}),
	}
	m := http.NewServeMux()
	m.Handle("/test", httptransport.NewServer(
		endpoint.AddEndpoint,
		decodeHTTPADDRequest,
		encodeHTTPGenericResponse,
		options...
	))

	return m
}

func decodeHTTPADDRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var (
		in  v2_service.TestRequest
		err error
	)
	in.A, err = strconv.Atoi(r.FormValue("a"))
	in.B, err = strconv.Atoi(r.FormValue("b"))
	if err != nil {
		return in, err
	}

	log.Println("开始请求数据")
	return in, nil
}

func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	log.Println("请求结束封装返回值")
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errWrapper struct {
	Error string `json:"error"`
}
