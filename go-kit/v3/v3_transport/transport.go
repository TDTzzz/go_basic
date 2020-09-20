package v3_transport

import (
	"context"
	"encoding/json"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"go_basic/go-kit/v3/utils"
	"go_basic/go-kit/v3/v3_endpoint"
	"go_basic/go-kit/v3/v3_service"
	"net/http"
	"strconv"
)

func NewHttpHandler(endpoint v3_endpoint.EndPointServer, log *zap.Logger) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(func(ctx context.Context, err error, w http.ResponseWriter) {
			log.Warn(fmt.Sprint(ctx.Value(v3_service.ContextReqUUid)), zap.Error(err))
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
		}),
		httptransport.ServerBefore(func(ctx context.Context, request *http.Request) context.Context {
			UUID := uuid.NewV5(uuid.Must(uuid.NewV4(), nil), "req_uuid").String()
			log.Debug("给请求添加uuid", zap.Any("UUID", UUID))
			ctx = context.WithValue(ctx, v3_service.ContextReqUUid, UUID)
			ctx = context.WithValue(ctx, utils.JWT_CONTEXT_KEY, request.Header.Get("Authorization"))
			log.Debug("把请求中的token发到Context中",
				zap.Any("Token", request.Header.Get("Authorization")))
			return ctx
		}),
	}
	m := http.NewServeMux()
	m.Handle("/sum", httptransport.NewServer(
		endpoint.AddEndPoint,
		decodeHTTPADDRequest,
		encodeHTTPGenericResponse,
		options...
	))
	m.Handle("/login", httptransport.NewServer(
		endpoint.LoginEndPoint,
		decodeHTTPLoginRequest,
		encodeHTTPGenericResponse,
		options...
	))
	return m
}

func decodeHTTPLoginRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var login v3_service.Login
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		return nil, err
	}
	utils.GetLogger().Debug(fmt.Sprint(ctx.Value(v3_service.ContextReqUUid)),
		zap.Any("开始解析请求数据", login))
	return login, nil
}

func decodeHTTPADDRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var (
		in  v3_service.Add
		err error
	)
	in.A, err = strconv.Atoi(r.FormValue("a"))
	in.B, err = strconv.Atoi(r.FormValue("b"))

	if err != nil {
		return in, err
	}
	utils.GetLogger().Debug(fmt.Sprint(ctx.Value(v3_service.ContextReqUUid)),
		zap.Any("开始解析请求数据", in))
	return in, nil
}

func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	utils.GetLogger().Debug(fmt.Sprint(ctx.Value(v3_service.ContextReqUUid)), zap.Any("请求结束封装返回值", response))
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorWrapper struct {
	Error string `json:"error"`
}
