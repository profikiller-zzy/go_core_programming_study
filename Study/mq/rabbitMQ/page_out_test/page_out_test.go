package page_out_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

const (
	rabbitmqURL = "amqp://itheima:123321@localhost:5672/"
	queueName   = "page_out_test_queue"
	messageSize = 1000000
)

func failOnError(t *testing.T, err error, msg string) {
	if err != nil {
		t.Fatalf("%s: %v", msg, err)
	}
}

func TestPageOutSend(t *testing.T) {
	start := time.Now()

	conn, err := amqp091.Dial(rabbitmqURL)
	failOnError(t, err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(t, err, "Failed to open a channel")
	defer ch.Close()

	// 声明持久化队列（durable=true）
	_, err = ch.QueueDeclare(
		queueName,
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,   // args
	)
	failOnError(t, err, "Failed to declare queue")

	// 生成一条 512 字节的消息体
	body := make([]byte, 512)
	for i := range body {
		body[i] = 'X'
	}

	t.Logf("Sending %d persistent messages to queue: %s", messageSize, queueName)

	for i := 1; i <= messageSize; i++ {
		err := ch.Publish(
			"",        // exchange
			queueName, // routing key
			false,     // mandatory
			false,     // immediate
			amqp091.Publishing{
				DeliveryMode: amqp091.Transient, // 是否是持久化消息
				ContentType:  "text/plain",
				Body:         []byte(fmt.Sprintf("Msg #%d %s", i, body)),
			})
		failOnError(t, err, fmt.Sprintf("Failed to publish message #%d", i))

		// 每发送 10000 条输出一次进度
		if i%10000 == 0 {
			t.Logf("Published %d messages", i)
		}
	}

	t.Logf("Finished sending %d messages in %s", messageSize, time.Since(start))
}
