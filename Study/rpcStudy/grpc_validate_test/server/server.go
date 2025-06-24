package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"

	pb "rpcStudy/grpc_validate_test/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.Person) (*pb.Reply, error) {
	log.Printf("Received request from: %s", req.Name)
	return &pb.Reply{Message: "Hello, " + req.Name}, nil
}

// Validator 为什么样定义这样一个接口呢？
// 因为在拦截器中，需要对传入的信息进行验证，如果只是断言成一个特定的类型，那么其他接口的请求就无法验证了
// 相反我们就是要调用它们的Validate()方法，这样就能实现对所有请求的验证
type Validator interface {
	Validate() error
}

func main() {
	// 1. 创建监听
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// 把metadata的验证写入到拦截器中，首先创建拦截器
	myInterceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// 如果这里仅仅断言成pb.Person，那么其他请求就无法验证了
		if r, ok := req.(Validator); ok { // 如果请求实现了Validator接口
			if err := r.Validate(); err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}
		return handler(ctx, req)
	}

	opt := grpc.UnaryInterceptor(myInterceptor)
	grpcServer := grpc.NewServer(opt)
	pb.RegisterGreeterServer(grpcServer, &server{})

	log.Println("gRPC server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
