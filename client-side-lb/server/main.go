package main

import (
	"context"
	"log"
	"net"

	"github.com/gofrs/uuid"
	pb "github.com/jun06t/grpc-sample/client-side-lb/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

var machine string

func init() {
	u := uuid.Must(uuid.NewV4())
	machine = u.String()
}

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message:   "Hello " + in.Name,
		MachineId: machine,
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
