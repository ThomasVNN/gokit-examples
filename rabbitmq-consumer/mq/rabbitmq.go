package mq

import (
	"encoding/base64"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/assembla/cony"
	"sync"
)

type Client struct {
	c     *cony.Client
	run   bool
	close chan bool
	wg    sync.WaitGroup
}

func Open(url string) *Client {
	c := &Client{
		c: cony.NewClient(
			cony.URL(url),
			cony.Backoff(cony.DefaultBackoff),
		),
		run:   true,
		close: make(chan bool, 0),
	}
	go c.runLoop()
	return c
}

func (c *Client) runLoop() {
	for c.run && c.c.Loop() {
		select {
		case err := <-c.c.Errors():
			if err == nil {
				continue
			}

			fmt.Println("mqclient error: %v", err)
		case <-c.close:
			fmt.Println("mqclient closed")
		}
	}
}

func (c *Client) ReadQueue(q *cony.Queue, f func(amqp.Delivery) error) {
	c.wg.Add(1)
	defer c.wg.Done()

	consumer := cony.NewConsumer(q)
	c.c.Consume(consumer)

	for c.run {
		select {
		case msg := <-consumer.Deliveries():
			if err := f(msg); err != nil {
				b64 := base64.StdEncoding.EncodeToString(msg.Body)
				fmt.Println("mqclient handler [%v] message: %v err: %v", q.Name, b64, err)
			}

			msg.Ack(false)
		case err := <-consumer.Errors():
			if err == nil {
				continue
			}

			fmt.Println("mqclient consumer [%v] err: %v", q.Name, err)
		case <-c.close:
		}
	}

	fmt.Println("mqclient read end queue: [%v]", q.Name)
}
