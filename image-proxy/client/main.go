package main

import (
	"context"
	"fmt"
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

	client := pb.NewConverterClient(conn)

	err = convert(client)
	if err != nil {
		log.Fatal(err)
	}
}

func convert(client pb.ConverterClient) error {
	file, err := os.Open("supercar.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stream, err := client.Convert(context.Background())
	if err != nil {
		return err
	}
	err = send(stream, file, "001")
	if err != nil {
		return err
	}

	err = receive(stream, "001")
	if err != nil {
		return err
	}

	return nil
}

func send(stream pb.Converter_ConvertClient, file *os.File, id string) error {
	meta := &pb.ConvertRequest{
		Value: &pb.ConvertRequest_Meta{Meta: &pb.Meta{Id: id, Type: "jpg", Quality: "90"}},
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

		data := &pb.ConvertRequest{
			Value: &pb.ConvertRequest_Chunk{Chunk: &pb.Chunk{Data: buf, Position: int64(n)}},
		}
		stream.Send(data)
	}

	err := stream.CloseSend()
	if err != nil {
		return err
	}

	return nil
}

func receive(stream pb.Converter_ConvertClient, id string) error {
	file, err := os.Create(fmt.Sprintf("%s.webp", id))
	if err != nil {
		return err
	}
	defer file.Close()

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		file.Write(resp.Data)
	}

	return nil
}
