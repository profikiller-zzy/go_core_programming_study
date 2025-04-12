package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/alibaba/sentinel-golang/util"
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
			Threshold:              10,
			StatIntervalInMs:       1000,
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject, // 超过请求频率的直接拒绝
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	ch := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func() {
			for {
				// 埋点逻辑，埋点资源名为 some-test
				e, b := sentinel.Entry("some-test")
				if b != nil {
					// 请求被拒绝，在此处进行处理
					time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
				} else {
					// 请求允许通过，此处编写业务逻辑
					fmt.Println(util.CurrentTimeMillis(), "Passed")
					time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)

					// 务必保证业务结束后调用 Exit
					e.Exit()
				}
			}
		}()
	}
	<-ch
}
