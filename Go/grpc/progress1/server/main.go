package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

type server struct {
	pb.UnimplementedGreeterServer
}

const (
	port = ":50051"
)

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Recieved: %v", req.GetName())
	return &pb.HelloReply{Message: "Hello" + req.Name}, nil
}

func main() {
	//创建 net.Listener 对象,监听 tcp 端口
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//创建 grpc server
	s := grpc.NewServer()
	//注册服务对应的实例
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listen to: %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
