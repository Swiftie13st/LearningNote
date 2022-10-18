package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc/pb/person"
	"io"
	"log"
	"strconv"
	"time"
)

func Simple(client person.SearchServiceClient) {
	res, err := client.Search(context.Background(), &person.PersonReq{Name: "Client1", Age: 18})
	if err != nil {
		log.Fatalf("could not send: %v", err)
	}
	fmt.Println(res)
}

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Close error: %v", err)
		}
	}(conn)

	client := person.NewSearchServiceClient(conn)

	c, err := client.SearchI(context.Background())
	if err != nil {
		log.Fatalf("%v.SearchI(_) = _, %v", client, err)
	}
	// 发送消息
	i := 0
	for {
		if i > 10 {
			res, err := c.CloseAndRecv()
			if err != nil {
				log.Fatalf("%v.CloseAndRecv() = _, %v", client, err)
				return
			}
			fmt.Println(res)
			break
		}
		time.Sleep(1 * time.Second)
		msg := &person.PersonReq{Name: "我是进来的信息" + strconv.Itoa(i)}
		if err := c.Send(msg); err != nil && err != io.EOF {
			log.Fatalf("%v.Send(%v) = %v", c, msg, err)
		}
		i++
	}
}
