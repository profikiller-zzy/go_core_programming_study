package main

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

type transactionListener struct {
}

func (tl *transactionListener) ExecuteLocalTransaction(msg *primitive.Message) primitive.LocalTransactionState {
	time.Sleep(time.Second * 2)
	// 假如说这里的本地事务突然挂掉了，什么都返回不了
	fmt.Printf("%v\n", time.Now().Format("2006-01-02 15:04:05"))
	//fmt.Println("假设这里的本地事务执行失败了，突然宕机")
	return primitive.CommitMessageState
}

func (tl *transactionListener) CheckLocalTransaction(msg *primitive.MessageExt) primitive.LocalTransactionState {
	fmt.Printf("%v", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("rocketMQ的事务回查\n")
	time.Sleep(time.Second * 2)
	// 假设这里因为代码错误，总是返回UnknowState，那么rocketMQ会一直回查（这个回查是有逻辑的）
	return primitive.UnknowState
}

func newTransactionListener() primitive.TransactionListener {
	return &transactionListener{}
}

func main() {
	p, _ := rocketmq.NewTransactionProducer(
		newTransactionListener(),
		producer.WithNameServer([]string{"8.155.39.153:9876"}),
		producer.WithRetry(1),
	)
	_ = p.Start()

	msg := primitive.NewMessage("mxshop_test", []byte("hello world"))
	res, err := p.SendMessageInTransaction(context.Background(), msg)
	if err != nil {
		fmt.Printf("send message error: %s\n", err)
	} else {
		fmt.Printf("send message success: result=%s\n", res.String())
	}

	time.Sleep(time.Hour)
	_ = p.Shutdown()
}
