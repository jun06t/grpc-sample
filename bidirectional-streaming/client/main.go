package main

import (
	"io"
	"log"

	pb "github.com/jun06t/grpc-sample/bidirectional-streaming/proto"
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

	client := pb.NewUppercaseServiceClient(conn)

	stream, err := client.Transform(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})

	go receive(done, stream)

	data := []string{"tokyo", "001", "Japan"}
	err = send(data, stream)
	if err != nil {
		log.Fatal(err)
	}

	<-done
}

func send(data []string, stream pb.UppercaseService_TransformClient) (err error) {
	for _, v := range data {
		log.Println("send message: ", v)
		err = stream.Send(&pb.UppercaseRequest{Message: v})
		if err != nil {
			return err
		}
	}
	err = stream.CloseSend()
	if err != nil {
		return err
	}

	return nil
}

func receive(done chan struct{}, stream pb.UppercaseService_TransformClient) {
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			// read done.
			close(done)
			return
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Println("received message: ", resp.Message)
	}
}
