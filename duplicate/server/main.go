package main

import (
	"context"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/any"
	pb "github.com/jun06t/grpc-sample/duplicate/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *any.Any) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello "}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
