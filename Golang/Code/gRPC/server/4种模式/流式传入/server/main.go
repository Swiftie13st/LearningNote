package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpc/pb/person"
	"io"
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
func (*personServer) SearchI(server person.SearchService_SearchIServer) error {
	for {
		req, err := server.Recv()
		fmt.Println(req)
		if err == io.EOF {
			err := server.SendAndClose(&person.PersonRes{Name: "完成了"})
			if err != nil {
				return err
			}
			break
		}
		if err != nil {
			return err
		}
	}
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
