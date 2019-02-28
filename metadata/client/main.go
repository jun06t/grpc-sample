package main

import (
	"context"
	"log"

	pb "github.com/jun06t/grpc-sample/metadata/proto"
	"github.com/lithammer/shortuuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	address = "localhost:8080"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	req := &pb.HelloRequest{
		Name: "alice",
	}
	ctx := setRequestID(context.Background())
	// Make RPC using the context with the metadata.
	var header, trailer metadata.MD
	resp, err := c.SayHello(ctx, req, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Reply: ", resp.Message)
	log.Printf("Header: %+v\n", header)
	log.Printf("Trailer: %+v\n", trailer)
}

func setRequestID(ctx context.Context) context.Context {
	id := shortuuid.New()
	return metadata.AppendToOutgoingContext(ctx, "x-request-id", id)
}
