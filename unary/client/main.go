package main

import (
	"log"

	pb "github.com/jun06t/grpc-sample/unary/proto"
	"golang.org/x/net/context"
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

	req := &pb.HelloRequest{
		Name: "alice",
	}
	resp, err := c.SayHello(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Reply: ", resp.Message)
}
