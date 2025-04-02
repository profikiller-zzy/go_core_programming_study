package main

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func main() {
	// push 模式的消费者 就是说如果服务器上有消息会主动推送给消费者
	c, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"http://8.155.39.153:9876"}),
		consumer.WithGroupName("mxshop"))
	if err != nil {
		panic(err)
	}

	// 订阅消息
	err = c.Subscribe("order_reback", consumer.MessageSelector{},
		func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for i := range msgs {
				fmt.Printf("received message: %s\n", msgs[i].Body)
			}
			return consumer.ConsumeSuccess, nil
		})
	if err != nil {
		fmt.Printf("订阅失败: %s\n", err)
	}
	_ = c.Start()
	time.Sleep(time.Hour)
	_ = c.Shutdown()
}
