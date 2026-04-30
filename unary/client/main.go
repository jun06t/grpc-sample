package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/jun06t/grpc-sample/unary/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:8080"
)

func main() {
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	fmt.Println("call api")
	req := &pb.HelloRequest{
		Name: "alice",
		Age:  10,
		Man:  true,
	}
	resp, err := c.SayHello(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Reply: ", resp.Message)
}
