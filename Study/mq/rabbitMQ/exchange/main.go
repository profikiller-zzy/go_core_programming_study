package main

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp091.Dial("amqp://itheima:123321@localhost:5672/")
	failOnError(err, "Failed to connect")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 1️⃣ 声明 fanout 类型的交换机
	err = ch.ExchangeDeclare(
		"fanout.exchange", // 交换机名
		"fanout",          // 类型
		true,              // durable
		false,             // auto-deleted
		false,             // internal
		false,             // no-wait
		nil,               // arguments
	)
	failOnError(err, "Failed to declare exchange")

	// 2️⃣ 声明一个持久队列
	q, err := ch.QueueDeclare(
		"fanout.queue.1", // 队列名
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		amqp091.Table{
			"x-queue-mode": "lazy", // 设置队列模式为 lazy
		}, // arguments
	)
	failOnError(err, "Failed to declare queue")

	// 3️⃣ 绑定队列到交换机（routing key 被忽略）
	err = ch.QueueBind(
		q.Name,            // 队列名
		"",                // routing key（fanout类型无效，但必须填）
		"fanout.exchange", // 交换机名
		false,
		nil,
	)
	failOnError(err, "Failed to bind queue")
}
