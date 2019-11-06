package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/hailinluo/gokit-examples/rabbitmq-consumer/mod"
	"github.com/hailinluo/gokit-examples/rabbitmq-consumer/service"
)


func MakeFxcmTradeEndpoint(svc service.MsgService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		msg := request.(*mod.Msg)
		if err := svc.HandleMsg(msg); err != nil {
			return nil, err
		}
		return nil, nil
	}
}
