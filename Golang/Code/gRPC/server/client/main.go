package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc/pb/person"
	"io"
	"log"
	"strconv"
	"sync"
	"time"
)

func simpleRpc(client person.SearchServiceClient) {
	res, err := client.Search(context.Background(), &person.PersonReq{Name: "Client1", Age: 18})
	if err != nil {
		log.Fatalf("could not send: %v", err)
	}
	fmt.Println(res)
}

func clientSideStreamingRpc(client person.SearchServiceClient) {
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

func serverSideStreamingRpc(client person.SearchServiceClient) {
	c, _ := client.SearchO(context.Background(), &person.PersonReq{Name: "Bruce"})
	for {
		res, err := c.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(res)
	}
}

func bidirectionalStreamingRPC(client person.SearchServiceClient) {
	c, err := client.SearchIO(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	// 流式发送
	go func() {
		for {
			time.Sleep(1 * time.Second)
			msg := &person.PersonReq{Name: "Bruce"}
			err := c.Send(msg)
			if err != nil {
				wg.Done()
				break
			}
			fmt.Printf("Client发送：%v\n", msg)
		}
	}()
	// 流式接收
	go func() {
		for {
			res, err := c.Recv()
			if err != nil {
				fmt.Println(err)
				wg.Done()
				break
			}
			fmt.Printf("Client接收：%v\n", res.Name)
		}
	}()
	wg.Wait()
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
	bidirectionalStreamingRPC(client)
}
