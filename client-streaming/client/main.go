package main

import (
	"io"
	"log"
	"os"

	pb "github.com/jun06t/grpc-sample/client-streaming/proto"
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

	client := pb.NewUploaderClient(conn)
	stream, err := client.Upload(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("supercar.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, 1024)
	for {
		_, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		stream.Send(&pb.Chunk{Data: buf})
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Status)
}
