// client.go
package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "testgrpc/pb" // 根据实际路径修改
)

const (
	address = "10.0.4.21:50051"
)

func main() {
	// 设置连接服务器的地址
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMyServiceClient(conn)

	// 构造请求消息
	req := &pb.Request{
		Message: "Hello from client",
	}

	// 发送请求给服务器
	res, err := c.Process(context.Background(), req)
	if err != nil {
		log.Fatalf("could not process request: %v", err)
	}

	// 处理服务器返回的响应
	log.Printf("Received response from server: %s", res.Result)
}
