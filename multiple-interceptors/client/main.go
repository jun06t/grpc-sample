package main

import (
	"context"
	"log"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"

	"github.com/jun06t/grpc-sample/multiple-interceptors/interceptor"
	pb "github.com/jun06t/grpc-sample/multiple-interceptors/proto"
)

const (
	address = "localhost:8080"
)

func main() {
	//run()
	runWithMiddleware()
}

func run() {
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(interceptor.ClientInterceptor("client-inter1-")),
		grpc.WithUnaryInterceptor(interceptor.ClientInterceptor("client-inter2-")),
		grpc.WithUnaryInterceptor(interceptor.ClientInterceptor("client-inter3-")),
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
}

func runWithMiddleware() {
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(
				interceptor.ClientInterceptor("client-inter1-"),
				interceptor.ClientInterceptor("client-inter2-"),
				interceptor.ClientInterceptor("client-inter3-"),
			),
		),
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
}
