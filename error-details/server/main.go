package main

import (
	"context"
	"log"
	"net"

	pb "github.com/jun06t/grpc-sample/error-details/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	port = ":8080"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	st := status.New(codes.InvalidArgument, "some error occurred")
	dt, _ := st.WithDetails(&pb.ErrorDetail{Code: pb.ErrorCode_INVALID_COUNTRY})
	return &pb.HelloReply{Message: "Hello " + in.Name}, dt.Err()
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
