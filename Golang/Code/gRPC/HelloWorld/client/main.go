package main

import (
	"fmt"
	hello_grpc "gRPC/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := hello_grpc.NewHelloGRPCClient(conn)
	req, err := client.SayHis(context.Background(), &hello_grpc.Req{Message: "我从客户端来"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(req.GetMessage())
}
