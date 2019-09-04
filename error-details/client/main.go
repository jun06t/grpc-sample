package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/jun06t/grpc-sample/error-details/proto"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
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
	resp, err := c.SayHello(context.Background(), req)
	if err != nil {
		handleError(err)
		log.Fatal(err)
	}
	log.Println("Reply: ", resp.Message)
}

func handleError(err error) {
	st, _ := status.FromError(err)

	for _, detail := range st.Details() {
		switch t := detail.(type) {
		case *errdetails.BadRequest:
			fmt.Println("handle BadRequest case")
		case *errdetails.QuotaFailure:
			fmt.Println("handle QuotaFailure case")
		case *pb.ErrorDetail:
			// handle original error code
			fmt.Println("error code:", t.Code)
		}
	}
}
