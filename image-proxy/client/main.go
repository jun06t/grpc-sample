package main

import (
	"context"
	"io"
	"log"
	"os"

	pb "github.com/jun06t/grpc-sample/image-proxy/proto"
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

	file, err := os.Open("supercar.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = upload(client, file)
	if err != nil {
		log.Fatal(err)
	}
}

func upload(client pb.UploaderClient, file *os.File) error {
	stream, err := client.Upload(context.Background())
	if err != nil {
		return err
	}
	meta := &pb.UploadRequest{
		Value: &pb.UploadRequest_Meta{Meta: &pb.Meta{Id: "001", Type: "", Quality: "90"}},
	}
	stream.Send(meta)

	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		data := &pb.UploadRequest{
			Value: &pb.UploadRequest_Chunk{Chunk: &pb.Chunk{Data: buf, Position: int64(n)}},
		}
		stream.Send(data)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	log.Println(resp.Status)
	return nil
}
