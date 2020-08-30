package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	pb "github.com/jun06t/grpc-sample/keepalive/proto"
)

const (
	address = "localhost:8080"
)

var kacp = keepalive.ClientParameters{
	Time:                10 * time.Second,
	Timeout:             5 * time.Second,
	PermitWithoutStream: true,
}

func main() {
	conn, err := grpc.Dial(
		address, grpc.WithInsecure(),
		grpc.WithKeepaliveParams(kacp),
	)
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
	select {} // Block forever; run with GODEBUG=http2debug=2 to observe ping frames and GOAWAYs due to idleness.
}
