package endpoint

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"go_basic/go-kit/v1/service"
	"strings"
)

type ArithmeticRequest struct {
	RequestType string `json:"request_type"`
	A           int    `json:"a"`
	B           int    `json:"b"`
}

type ArithmeticResponse struct {
	Result int   `json:"result"`
	Error  error `json:"error"`
}

func MakeArithmeticEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ArithmeticRequest)

		var (
			res, a, b int
			calErr    error
		)
		a = req.A
		b = req.B

		if strings.EqualFold(req.RequestType, "Add") {
			res = svc.Add(a, b)
		} else if strings.EqualFold(req.RequestType, "Subtract") {
			res = svc.Subtract(a, b)
		} else {
			//这里应该定个error
			return nil, errors.New("算数方法错了")
		}

		return ArithmeticResponse{
			Result: res,
			Error:  calErr,
		}, nil
	}
}
