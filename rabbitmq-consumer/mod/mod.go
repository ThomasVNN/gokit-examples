package mod

import "github.com/streadway/amqp"

type Msg struct {
	Msg   string
	Acknowledger amqp.Acknowledger
	DeliveryTag uint64
}
