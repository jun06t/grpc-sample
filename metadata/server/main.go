package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/jun06t/grpc-sample/metadata/interceptor"
	pb "github.com/jun06t/grpc-sample/metadata/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	port = ":8080"
)

func main() {
	//run()
	runWithInterceptor()
}

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	defer func() {
		trailer := metadata.Pairs("timestamp", time.Now().Format(time.StampNano))
		grpc.SetTrailer(ctx, trailer)
	}()
	header := metadata.Pairs("timestamp", time.Now().Format(time.StampNano))
	grpc.SetHeader(ctx, header)

	// Read metadata from client.
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.DataLoss, "UnarySayHello: failed to get metadata")
	}
	if t, ok := md["x-request-id"]; ok {
		fmt.Printf("request id from metadata:\n")
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	}

	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

type serverWithInterceptor struct{}

func (s *serverWithInterceptor) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func run() {
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

func runWithInterceptor() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.ServerInterceptor),
	)
	pb.RegisterGreeterServer(s, &serverWithInterceptor{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
