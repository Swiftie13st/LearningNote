package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc/pb/person"
	"log"
	"net"
	"net/http"
	"sync"
)

type personServer struct {
	person.UnimplementedSearchServiceServer
}

// Search 基本模式
func (*personServer) Search(ctx context.Context, req *person.PersonReq) (*person.PersonRes, error) {
	name := req.GetName()
	res := &person.PersonRes{
		Name: "我收到了" + name + "的信息",
	}
	return res, nil
}

func registerGateway(wg *sync.WaitGroup) {
	conn, err := grpc.DialContext(context.Background(), "127.0.0.1:8888", grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}

	mux := runtime.NewServeMux() // 一个对外开放的mux

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: mux,
	}

	err = person.RegisterSearchServiceHandler(context.Background(), mux, conn)
	if err != nil {
		return
	}

	err = gwServer.ListenAndServe()
	if err != nil {
		return
	}
	wg.Done()

}

func registerGRPC(wg *sync.WaitGroup) {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	person.RegisterSearchServiceServer(s, &personServer{})

	log.Printf("server listening at %v", l.Addr())
	// 建立监听
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go registerGateway(&wg)
	go registerGRPC(&wg)
	wg.Wait()
}
