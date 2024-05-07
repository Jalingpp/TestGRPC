package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "testgrpc/pb" // 根据实际路径修改
)

type server struct {
	pb.UnimplementedMyServiceServer // 嵌入未实现的 gRPC 服务器结构体
}

func (s *server) Process(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Printf("Received message from client: %s", req.Message)
	// 在这里处理请求，并生成响应
	return &pb.Response{Result: "Processed: " + req.Message}, nil
}

func main() {
	lis, err := net.Listen("tcp", "10.0.4.21:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMyServiceServer(s, &server{})
	log.Println("Server listening on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
