package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpc/pb/person"
	"log"
	"net"
)

type personServer struct {
	person.UnimplementedSearchServiceServer
}

func (*personServer) Search(ctx context.Context, req *person.PersonReq) (*person.PersonRes, error) {
	name := req.GetName()
	res := &person.PersonRes{
		Name: "我收到了" + name + "的信息",
	}
	return res, nil
}
func (*personServer) SearchI(person.SearchService_SearchIServer) error {
	return nil
}
func (*personServer) SearchO(*person.PersonReq, person.SearchService_SearchOServer) error {
	return nil
}
func (*personServer) SearchIO(person.SearchService_SearchIOServer) error {
	return nil
}

func main() {
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
}
