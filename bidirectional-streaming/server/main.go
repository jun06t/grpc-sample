package main

import (
	"io"
	"log"
	"net"
	"strings"

	pb "github.com/jun06t/grpc-sample/bidirectional-streaming/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type server struct{}

func (s *server) Transform(stream pb.UppercaseService_TransformServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		resp := &pb.UppercaseResponse{
			Message: strings.ToUpper(in.Message),
		}
		err = stream.Send(resp)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterUppercaseServiceServer(s, new(server))
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
