package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/pb/person"
	"io"
	"log"
	"net"
	"strconv"
	"time"
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

// SearchI 流式传入
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

// SearchO 流式传出
func (*personServer) SearchO(req *person.PersonReq, server person.SearchService_SearchOServer) error {
	name := req.Name
	i := 0
	for {
		if i > 10 {
			break
		}
		time.Sleep(1 * time.Second)
		err := server.Send(&person.PersonRes{Name: "我拿到了" + name + strconv.Itoa(i)})
		if err != nil {
			return err
		}
		i++
	}
	return nil
}

// SearchIO 流式出入
func (*personServer) SearchIO(server person.SearchService_SearchIOServer) error {
	i := 0
	// 流式接收，通过channel传输信息
	str := make(chan string)
	go func() {
		for {
			i++
			req, _ := server.Recv()
			fmt.Printf("Server接收：%v\n", req)
			if i > 10 {
				str <- "结束"
				break
			}

			str <- req.Name + strconv.Itoa(i)
		}
	}()

	for {
		s := <-str
		fmt.Printf("Server发送：%v\n", s)
		if s == "结束" {
			// 结束
			server.Send(&person.PersonRes{Name: s})
			break
		}
		server.Send(&person.PersonRes{Name: s})

	}
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
