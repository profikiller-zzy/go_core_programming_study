package main

import (
	"context"
	"log"
	"net"
	pb "rpcStudy/grpc_error_test/proto"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	// 模拟网络拥塞
	time.Sleep(3 * time.Second)
	log.Printf("Received request from: %s", req.Name)
	return &pb.HelloReply{Message: "Hello, " + req.Name}, nil
	//return nil, status.New(codes.InvalidArgument, "非法名字").Err()
}

func main() {
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
