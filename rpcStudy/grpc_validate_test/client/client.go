package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "rpcStudy/grpc_validate_test/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	resp, err := client.SayHello(context.Background(), &pb.Person{
		Id:    999,
		Name:  "Alice",
		Email: "3548361574@qq.com",
	})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	fmt.Println("Server Response:", resp.Message)
}
