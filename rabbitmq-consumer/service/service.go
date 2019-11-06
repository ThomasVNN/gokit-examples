package service

import (
	"fmt"
	"github.com/hailinluo/gokit-examples/rabbitmq-consumer/mod"
	"time"
)

type MsgService interface {
	HandleMsg(*mod.Msg) error
}

type MsgHandler struct{}

func (MsgHandler) HandleMsg(msg *mod.Msg) error {
	for i := 0; i < 30; i++ {
		time.Sleep(time.Second)
		fmt.Println("doing...")
	}
	msg.Acknowledger.Ack(msg.DeliveryTag, false)
	fmt.Println("handle msg finished, msg: ", msg.Msg)
	return nil
}

