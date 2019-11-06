package transport

import (
	"context"
	"github.com/hailinluo/gokit-examples/rabbitmq-consumer/mod"
	"github.com/streadway/amqp"
)

func DecodeFxcmTradeRequest(_ context.Context, deliv *amqp.Delivery) (interface{}, error) {
	request := new(mod.Msg)
	request.Msg = string(deliv.Body)
	request.Acknowledger = deliv.Acknowledger
	request.DeliveryTag = deliv.DeliveryTag
	return request, nil
}
