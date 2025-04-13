package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
)

func main() {
	// 务必先进行初始化
	err := sentinel.InitDefault()
	if err != nil {
		log.Fatal(err)
	}

	// 配置一条限流规则
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "some-test",
			Threshold:              2000,
			WarmUpPeriodSec:        10, // 预热的时间长度，也就是说在10秒之内慢慢达到设置的流量最大频率
			StatIntervalInMs:       1000,
			TokenCalculateStrategy: flow.WarmUp, // 冷启动策略
			ControlBehavior:        flow.Reject, // 超过请求频率的直接拒绝
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	var (
		count int
		lock  int
		pass  int
	)
	ch := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func() {
			for {
				// 埋点逻辑，埋点资源名为 some-test
				e, b := sentinel.Entry("some-test")
				count++
				if b != nil {
					// 请求被拒绝，在此处进行处理
					lock++
					time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
				} else {
					pass++
					// 请求允许通过，此处编写业务逻辑
					time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
					// 务必保证业务结束后调用 Exit
					e.Exit()
				}
			}
		}()
	}

	go func() {
		for {
			curCount, curLock, curPass := count, lock, pass
			time.Sleep(1 * time.Second)
			fmt.Printf("这一秒内的流量，总：%v, 通过的流量：%v, 被拒绝的流量：%v\n", count-curCount, pass-curPass, lock-curLock)
		}
	}()
	<-ch
}
