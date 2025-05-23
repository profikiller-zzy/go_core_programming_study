package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"github.com/alibaba/sentinel-golang/core/config"
	"github.com/alibaba/sentinel-golang/logging"
)

type stateChangeTestListener struct {
}

func (s *stateChangeTestListener) OnTransformToClosed(prev circuitbreaker.State, rule circuitbreaker.Rule) {
	fmt.Printf("rule.steategy: %+v, From %s to Closed, time: %v\n", rule.Strategy, prev.String(), time.Now().Format("2006-01-02 15:04:05"))
}

func (s *stateChangeTestListener) OnTransformToOpen(prev circuitbreaker.State, rule circuitbreaker.Rule, snapshot interface{}) {
	fmt.Printf("rule.steategy: %+v, From %s to Open, snapshot: %.2f, time: %v\n", rule.Strategy, prev.String(), snapshot, time.Now().Format("2006-01-02 15:04:05"))
}

func (s *stateChangeTestListener) OnTransformToHalfOpen(prev circuitbreaker.State, rule circuitbreaker.Rule) {
	fmt.Printf("rule.steategy: %+v, From %s to Half-Open, time: %v\n", rule.Strategy, prev.String(), time.Now().Format("2006-01-02 15:04:05"))
}

func main() {
	conf := config.NewDefaultConfig()
	// for testing, logging output to console
	conf.Sentinel.Log.Logger = logging.NewConsoleLogger()
	err := sentinel.InitWithConfig(conf)
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan struct{})
	// Register a state change listener so that we could observer the state change of the internal circuit breaker.
	circuitbreaker.RegisterStateChangeListeners(&stateChangeTestListener{})

	_, err = circuitbreaker.LoadRules([]*circuitbreaker.Rule{
		{
			Resource:                     "abc",
			Strategy:                     circuitbreaker.SlowRequestRatio,
			RetryTimeoutMs:               3000, // 3s之后尝试恢复
			MinRequestAmount:             10,   // 静默数量
			StatIntervalMs:               5000, // 统计的时间窗口长度
			StatSlidingWindowBucketCount: 10,
			MaxAllowedRtMs:               50,  // 请求响应时间大于该值视为慢请求
			Threshold:                    0.5, // 慢调用比例的阈值
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	logging.Info("[CircuitBreaker SlowRtRatio] Sentinel Go circuit breaking demo is running. You may see the pass/block metric in the metric log.")
	go func() {
		for {
			e, b := sentinel.Entry("abc")
			if b != nil {
				// g1 blocked
				time.Sleep(time.Duration(rand.Uint64()%20) * time.Millisecond)
			} else {
				if rand.Uint64()%20 > 9 {
					// Record current invocation as error.
					sentinel.TraceError(e, errors.New("biz error"))
				}
				// g1 passed
				time.Sleep(time.Duration(rand.Uint64()%80+10) * time.Millisecond)
				e.Exit()
			}
		}
	}()
	go func() {
		for {
			e, b := sentinel.Entry("abc")
			if b != nil {
				// g2 blocked
				time.Sleep(time.Duration(rand.Uint64()%20) * time.Millisecond)
			} else {
				// g2 passed
				time.Sleep(time.Duration(rand.Uint64()%80+10) * time.Millisecond)
				e.Exit()
			}
		}
	}()
	<-ch
}
