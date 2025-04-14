package main

import (
	"context"
	"log"
	"net"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"google.golang.org/grpc"

	pb "go_core_programming/rpcStudy/metadata_test/proto"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received request from: %s", req.Name)
	return &pb.HelloReply{Message: "Hello, " + req.Name}, nil
}

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

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &server{})

	log.Println("gRPC server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
