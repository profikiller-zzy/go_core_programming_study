package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func main() {
	p, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{"http://8.155.39.153:9876"}),
		producer.WithRetry(2),
	)
	defer func(p rocketmq.Producer) {
		err := p.Shutdown()
		if err != nil {
			fmt.Printf("shut down producer error: %s\n", err)
		}
	}(p)
	if err != nil {
		panic(err)
	}

	if err := p.Start(); err != nil {
		panic(err)
	}

	res, err := p.SendSync(context.Background(), primitive.NewMessage("mxshop_test", []byte("hello this is rocketMQ go client1")))
	if err != nil {
		fmt.Printf("send message error: %s\n", err)
	} else {
		fmt.Printf("send message success: result=%s\n", res.String())
	}
}
