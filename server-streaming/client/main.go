package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/jun06t/grpc-sample/server-streaming/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
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

	client := pb.NewFeederClient(conn)

	stream, err := client.GetNewFeed(context.Background(), new(pb.Empty))
	if err != nil {
		log.Fatal(err)
	}
	for {
		article, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(article)
	}
}
