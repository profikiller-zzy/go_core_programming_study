package main

import (
	"context"
	"go.opentelemetry.io/otel/trace"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
)

func initTracer() func(context.Context) error {
	// 使用 WithCollectorEndpoint 指定 Jaeger Collector 的 HTTP 接口地址
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(
		jaeger.WithEndpoint("http://8.155.39.153:14268/api/traces"), // Jaeger HTTP 接口
	))
	if err != nil {
		log.Fatal(err)
	}

	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("example-service"),
		)),
	)

	// 将刚刚配置的jaeger服务端点注册到全局的TracerProvider
	otel.SetTracerProvider(tp)

	return tp.Shutdown
}

func main() {
	shutdown := initTracer()
	defer shutdown(context.Background())

	// 初始化tracer
	tracer := otel.Tracer("mxshop-example-tracer")

	// 创建根 Context
	ctx := context.Background()

	// 第一个 Span（根 Span）
	ctx, span := tracer.Start(ctx, "function1")
	time.Sleep(500 * time.Millisecond)
	span.End()

	// 第二个 Span（FollowsFrom 关系）
	options := []trace.SpanStartOption{
		trace.WithLinks(trace.LinkFromContext(ctx)),
	}
	_, span2 := tracer.Start(ctx, "function2", options...)
	time.Sleep(500 * time.Millisecond)
	span2.End()
}
