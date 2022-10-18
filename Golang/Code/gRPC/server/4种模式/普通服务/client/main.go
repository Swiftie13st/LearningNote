package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc/pb/person"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := person.NewSearchServiceClient(conn)
	res, err := client.Search(context.Background(), &person.PersonReq{Name: "Client1", Age: 18})
	if err != nil {
		log.Fatalf("could not send: %v", err)
	}
	fmt.Println(res)
}
