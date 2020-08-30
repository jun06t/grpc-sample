package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	pb "github.com/jun06t/grpc-sample/keepalive/proto"
)

const (
	port = ":8080"
)

var kaep = keepalive.EnforcementPolicy{
	MinTime:             10 * time.Second,
	PermitWithoutStream: true,
}

var kasp = keepalive.ServerParameters{
	MaxConnectionIdle:     30 * time.Second,
	MaxConnectionAge:      30 * time.Second,
	MaxConnectionAgeGrace: 5 * time.Second,
	Time:                  12 * time.Second,
	Timeout:               3 * time.Second,
}

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	time.Sleep(30 * time.Second)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer(
		grpc.KeepaliveEnforcementPolicy(kaep),
		grpc.KeepaliveParams(kasp),
	)
	pb.RegisterGreeterServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
