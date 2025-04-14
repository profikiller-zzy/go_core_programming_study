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
			semconv.SchemaURL, // 配置资源属性
			semconv.ServiceNameKey.String("example-service"), // 配置服务名称
		)),
	)

	// 将刚刚配置的jaeger服务端点注册到全局的 TracerProvider
	otel.SetTracerProvider(tp)

	return tp.Shutdown
}

func function2(ctx context.Context, tracer trace.Tracer) {
	// function2 span 继承自 function1 span
	ctx, span := tracer.Start(ctx, "function2")
	defer span.End()

	time.Sleep(500 * time.Millisecond)
}

func main() {
	shutdown := initTracer()
	defer shutdown(context.Background())

	tracer := otel.Tracer("mxshop-example-tracer")
	mainCtx := context.Background()

	// 根 Span
	mainCtx, mainSpan := tracer.Start(mainCtx, "main")
	defer mainSpan.End()

	// 创建子 Span
	_, span1 := tracer.Start(mainCtx, "function 1")
	time.Sleep(500 * time.Millisecond)
	span1.End()

	function2(mainCtx, tracer) // 调用子函数并传入 ctx
}
