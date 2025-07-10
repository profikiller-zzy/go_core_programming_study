package main

import (
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	go consumer("consumer_fast", 50, "[FAST]")
	go consumer("consumer_slow", 5, "[SLOW]")

	select {} // 阻塞主线程
}

func consumer(consumerTag string, qos int, label string) {
	conn, err := amqp091.Dial("amqp://itheima:123321@localhost:5672/")
	if err != nil {
		log.Fatalf("%s Failed to connect: %v", label, err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("%s Failed to open channel: %v", label, err)
	}

	err = ch.Qos(qos, 0, false)
	if err != nil {
		log.Fatalf("%s Failed to set QoS: %v", label, err)
	}

	msgs, err := ch.Consume("work.queue", consumerTag, false, false, false, false, nil)
	if err != nil {
		log.Fatalf("%s Failed to register consumer: %v", label, err)
	}

	log.Printf("%s Waiting for messages...", label)

	ticker := time.NewTicker(time.Second)

	buffer := make([]amqp091.Delivery, 0, qos)

	for {
		select {
		case d, ok := <-msgs:
			if !ok {
				log.Printf("%s Channel closed", label)
				return
			}
			buffer = append(buffer, d)
			if len(buffer) >= qos {
				for _, m := range buffer {
					log.Printf("%s Received: %s", label, m.Body)
					m.Ack(false)
				}
				buffer = buffer[:0]
			}
		case <-ticker.C:
			for _, m := range buffer {
				log.Printf("%s Received (tick): %s", label, m.Body)
				m.Ack(false)
			}
			buffer = buffer[:0]
		}
	}
}
