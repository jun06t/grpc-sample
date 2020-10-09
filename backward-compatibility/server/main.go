package main

import (
	"context"
	"log"
	"net"

	pb "github.com/jun06t/grpc-sample/backward-compatibility/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "Hello " + in.Name,
		Renamed: &pb.Renamed{
			Amount: 100,
			Name:   "foobar",
			Term: &pb.Renamed_Term{
				StartAt: 1234567890,
				EndAt:   9876543210,
			},
		},
	}, nil
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
