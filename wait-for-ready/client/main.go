package main

import (
	"context"
	"log"
	"time"

	pb "github.com/jun06t/grpc-sample/wait-for-ready/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	for {
		hello(c)
		time.Sleep(1 * time.Second)
	}
}

func hello(c pb.GreeterClient) {
	req := &pb.HelloRequest{
		Name: "alice",
		Age:  10,
		Man:  true,
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	resp, err := c.SayHello(ctx, req, grpc.WaitForReady(true))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Reply: ", resp.Message)
}
