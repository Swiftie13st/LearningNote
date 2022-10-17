package main

import (
	"context"
	"fmt"
	hello_grpc "gRPC/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	hello_grpc.UnimplementedHelloGRPCServer
}

func (receiver *server) SayHis(ctx context.Context, req *hello_grpc.Req) (res *hello_grpc.Res, err error) {
	fmt.Println(req.GetMessage())
	return &hello_grpc.Res{Message: "服务端返回的grpc内容"}, nil
}

func main() {
	// 注册服务
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hello_grpc.RegisterHelloGRPCServer(s, &server{})
	log.Printf("server listening at %v", l.Addr())
	// 建立监听
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
