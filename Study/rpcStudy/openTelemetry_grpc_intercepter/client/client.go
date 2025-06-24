package main

import (
	"context"
	"fmt"
	"go_core_programming/Study/rpcStudy/metadata_test/proto"
	"log"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"google.golang.org/grpc"
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

func main() {
	shutdown := initTracer()
	defer shutdown(context.Background())

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithStatsHandler(otelgrpc.NewClientHandler()))

	// 2. 创建连接，设置拦截器
	conn, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := proto.NewGreeterClient(conn)

	resp, err := client.SayHello(context.Background(), &proto.HelloRequest{Name: "Alice"})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	fmt.Println("Server Response:", resp.Message)
}
