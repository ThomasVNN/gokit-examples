package main

import (
	"fmt"
	"github.com/assembla/cony"
	amqptransport "github.com/go-kit/kit/transport/amqp"
	"github.com/hailinluo/gokit-examples/rabbitmq-consumer/endpoint"
	"github.com/hailinluo/gokit-examples/rabbitmq-consumer/mq"
	"github.com/hailinluo/gokit-examples/rabbitmq-consumer/service"
	"github.com/hailinluo/gokit-examples/rabbitmq-consumer/transport"
	"github.com/streadway/amqp"
)

func main() {
	rabbit := mq.Open("amqp://guest:guest@127.0.0.1:5672")

	srv := service.MsgHandler{}
	subscriber := amqptransport.NewSubscriber(
		endpoint.MakeFxcmTradeEndpoint(srv),
		transport.DecodeFxcmTradeRequest,
		amqptransport.EncodeNopResponse, // No response encode
		// No Subscribe Response
		amqptransport.SubscriberResponsePublisher(amqptransport.NopResponsePublisher),
	)

	rabbit.ReadQueue(&cony.Queue{Name: "signal.one"}, func(delivery amqp.Delivery) error {
		fmt.Println("receive message from fxcmTradeSignalQueue")
		subscriber.ServeDelivery(nil)(&delivery)
		return nil
	})
}
