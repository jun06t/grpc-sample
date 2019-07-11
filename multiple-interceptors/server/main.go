package main

import (
	"context"
	"fmt"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"

	"github.com/jun06t/grpc-sample/multiple-interceptors/interceptor"
	pb "github.com/jun06t/grpc-sample/multiple-interceptors/proto"
)

const (
	port = ":8080"
)

func main() {
	// run()
	runWithMiddleware()
}

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("handler")
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func run() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.ServerInterceptor("server-inter1-")),
		grpc.UnaryInterceptor(interceptor.ServerInterceptor("server-inter2-")),
		grpc.UnaryInterceptor(interceptor.ServerInterceptor("server-inter3-")),
	)
	pb.RegisterGreeterServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}

func runWithMiddleware() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			interceptor.ServerInterceptor("server-inter1-"),
			interceptor.ServerInterceptor("server-inter2-"),
			interceptor.ServerInterceptor("server-inter3-"),
		),
	)
	pb.RegisterGreeterServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
