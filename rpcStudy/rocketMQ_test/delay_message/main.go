package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"time"
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

	msg := primitive.NewMessage("mxshop_test", []byte("hello this is rocketMQ go client1, delay message"))
	// 1s 5s 10s 30s 1m 2m 3m 4m 5m 6m 7m 8m 9m 10m 20m 30m 1h 2h
	// 1  2   3   4   5  6  7  8  9 10 11 12 13 14  15  16 17 18
	msg.WithDelayTimeLevel(2) // 延迟三秒发送
	res, err := p.SendSync(context.Background(), msg)
	if err != nil { // 并不能使用err != nil来判断是否发送成功，需要检查res.Status
		fmt.Printf("send message error: %s\n", err)
	} else {
		fmt.Printf("send message success: result=%s\n", res.String())
	}

	time.Sleep(time.Second * 5)
	if res.Status == primitive.SendOK {
		fmt.Println("send success")
	} else {
		fmt.Println("send failed")
	}
}
