package main

import (
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type ProducerConfig struct {
	MaxRetries    int
	RetryInterval time.Duration
	Timeout       time.Duration
}

type Producer struct {
	config ProducerConfig
	conn   *amqp.Connection
	ch     *amqp.Channel
}

func NewProducer(config ProducerConfig) *Producer {
	return &Producer{
		config: config,
	}
}

func (p *Producer) Connect() error {
	var err error
	for attempt := 1; attempt <= p.config.MaxRetries; attempt++ {
		log.Printf("尝试连接RabbitMQ (第%d次)...", attempt)
		
		// 设置连接超时
		config := amqp.Config{
			Dial: amqp.DefaultDial(p.config.Timeout),
		}
		
		p.conn, err = amqp.DialConfig("amqp://itheima:123321@localhost:5672/", config)
		if err != nil {
			log.Printf("连接失败 (第%d次): %v", attempt, err)
			if attempt < p.config.MaxRetries {
				log.Printf("等待%v后重试...", p.config.RetryInterval)
				time.Sleep(p.config.RetryInterval)
				continue
			}
			return fmt.Errorf("连接失败，已重试%d次: %v", p.config.MaxRetries, err)
		}
		
		p.ch, err = p.conn.Channel()
		if err != nil {
			log.Printf("创建通道失败 (第%d次): %v", attempt, err)
			p.conn.Close()
			if attempt < p.config.MaxRetries {
				time.Sleep(p.config.RetryInterval)
				continue
			}
			return fmt.Errorf("创建通道失败，已重试%d次: %v", p.config.MaxRetries, err)
		}
		
		log.Printf("连接成功！")
		return nil
	}
	return err
}

func (p *Producer) DeclareQueue(queueName string) error {
	_, err := p.ch.QueueDeclare(queueName, true, false, false, false, nil)
	return err
}

func (p *Producer) PublishWithRetry(queueName, message string) error {
	for attempt := 1; attempt <= p.config.MaxRetries; attempt++ {
		log.Printf("尝试发送消息 (第%d次): %s", attempt, message)
		
		err := p.ch.Publish(
			"",        // exchange
			queueName, // routing key
			false,     // mandatory
			false,     // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(message),
				Timestamp:   time.Now(),
			},
		)
		
		if err != nil {
			log.Printf("发送消息失败 (第%d次): %v", attempt, err)
			
			// 检查是否是连接问题
			if p.conn.IsClosed() || p.ch.IsClosed() {
				log.Printf("检测到连接断开，尝试重新连接...")
				if reconnectErr := p.Connect(); reconnectErr != nil {
					log.Printf("重新连接失败: %v", reconnectErr)
					if attempt < p.config.MaxRetries {
						time.Sleep(p.config.RetryInterval)
						continue
					}
					return fmt.Errorf("重新连接失败，已重试%d次: %v", p.config.MaxRetries, reconnectErr)
				}
				
				// 重新声明队列
				if declareErr := p.DeclareQueue(queueName); declareErr != nil {
					log.Printf("重新声明队列失败: %v", declareErr)
				}
			}
			
			if attempt < p.config.MaxRetries {
				log.Printf("等待%v后重试...", p.config.RetryInterval)
				time.Sleep(p.config.RetryInterval)
				continue
			}
			return fmt.Errorf("发送消息失败，已重试%d次: %v", p.config.MaxRetries, err)
		}
		
		log.Printf("消息发送成功: %s", message)
		return nil
	}
	return nil
}

func (p *Producer) Close() {
	if p.ch != nil {
		p.ch.Close()
	}
	if p.conn != nil {
		p.conn.Close()
	}
}

// 模拟网络故障的函数
func simulateNetworkFailure() {
	log.Println("=== 模拟网络故障场景 ===")
	
	// 配置：最多重试5次，每次间隔2秒，连接超时3秒
	config := ProducerConfig{
		MaxRetries:    5,
		RetryInterval: 2 * time.Second,
		Timeout:       3 * time.Second,
	}
	
	producer := NewProducer(config)
	defer producer.Close()
	
	// 尝试连接（可能失败）
	if err := producer.Connect(); err != nil {
		log.Fatalf("最终连接失败: %v", err)
	}
	
	// 声明队列
	if err := producer.DeclareQueue("retry.queue"); err != nil {
		log.Fatalf("声明队列失败: %v", err)
	}
	
	// 发送消息（带重试）
	for i := 1; i <= 10; i++ {
		message := fmt.Sprintf("重试测试消息 #%d - %s", i, time.Now().Format("15:04:05"))
		
		if err := producer.PublishWithRetry("retry.queue", message); err != nil {
			log.Printf("消息 #%d 最终发送失败: %v", i, err)
		} else {
			log.Printf("消息 #%d 发送成功", i)
		}
		
		// 间隔发送
		time.Sleep(1 * time.Second)
	}
}

func main() {
	log.Println("启动带重试机制的RabbitMQ生产者")
	log.Println("提示：请确保RabbitMQ服务正在运行，或者故意停止服务来测试重试机制")
	
	simulateNetworkFailure()
	
	log.Println("测试完成")
}