package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/hailinluo/gokit-examples/strings-srv/service"
	"github.com/hailinluo/gokit-examples/strings-srv/transport"
)

func makeUppercaseEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.UpercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return transport.UppercaseResponse{v, err.Error()}, nil
		}
		return transport.UppercaseResponse{v, ""}, nil
	}
}

func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.CountRequest)
		v := svc.Count(req.S)
		return transport.CountResponse{v}, nil
	}
}
