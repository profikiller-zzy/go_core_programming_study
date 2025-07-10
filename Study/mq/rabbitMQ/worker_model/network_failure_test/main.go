package main

import (
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func testConnectionFailure() {
	log.Println("=== 测试连接失败重试 ===")
	
	// 故意使用错误的地址来模拟网络故障
	wrongAddresses := []string{
		"amqp://itheima:123321@192.168.1.999:5672/", // 不存在的IP
		"amqp://itheima:123321@localhost:9999/",     // 错误的端口
		"amqp://itheima:123321@localhost:5672/",     // 正确的地址（最后尝试）
	}
	
	for _, addr := range wrongAddresses {
		log.Printf("尝试连接: %s", addr)
		
		config := amqp.Config{
			Dial: amqp.DefaultDial(3 * time.Second), // 3秒超时
		}
		
		start := time.Now()
		conn, err := amqp.DialConfig(addr, config)
		duration := time.Since(start)
		
		if err != nil {
			log.Printf("连接失败 (耗时: %v): %v", duration, err)
			time.Sleep(2 * time.Second) // 重试间隔
			continue
		}
		
		log.Printf("连接成功 (耗时: %v)", duration)
		conn.Close()
		break
	}
}

func testPublishFailure() {
	log.Println("\n=== 测试发布消息失败重试 ===")
	
	conn, err := amqp.Dial("amqp://itheima:123321@localhost:5672/")
	if err != nil {
		log.Printf("无法连接到RabbitMQ进行发布测试: %v", err)
		return
	}
	defer conn.Close()
	
	ch, err := conn.Channel()
	if err != nil {
		log.Printf("无法创建通道: %v", err)
		return
	}
	defer ch.Close()
	
	// 声明队列
	_, err = ch.QueueDeclare("test.queue", true, false, false, false, nil)
	if err != nil {
		log.Printf("声明队列失败: %v", err)
		return
	}
	
	log.Println("开始发送消息...")
	log.Println("提示：在发送过程中，你可以停止RabbitMQ服务来模拟网络故障")
	
	for i := 1; i <= 20; i++ {
		message := fmt.Sprintf("测试消息 #%d - %s", i, time.Now().Format("15:04:05"))
		
		err := ch.Publish(
			"",
			"test.queue",
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(message),
			},
		)
		
		if err != nil {
			log.Printf("发送消息 #%d 失败: %v", i, err)
			
			// 检查连接状态
			if conn.IsClosed() {
				log.Println("检测到连接已关闭")
				break
			}
		} else {
			log.Printf("消息 #%d 发送成功", i)
		}
		
		time.Sleep(1 * time.Second)
	}
}

func main() {
	log.Println("RabbitMQ网络故障测试")
	log.Println("==================")
	
	// 测试连接失败
	testConnectionFailure()
	
	// 测试发布失败
	testPublishFailure()
	
	log.Println("\n测试完成")
}