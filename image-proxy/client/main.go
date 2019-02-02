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

	file, err := os.Open("testimage.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = convert(client, file, "001")
	if err != nil {
		log.Fatal(err)
	}
}

func convert(client pb.ConverterClient, src io.Reader, name string) error {
	stream, err := client.Convert(context.Background())
	if err != nil {
		return err
	}
	err = send(stream, src, name)
	if err != nil {
		return err
	}

	dst, err := os.Create(fmt.Sprintf("%s.webp", name))
	if err != nil {
		return err
	}
	defer dst.Close()

	err = receive(stream, dst)
	if err != nil {
		return err
	}

	return nil
}

const (
	bufSize = 1024
)

func send(stream pb.Converter_ConvertClient, src io.Reader, id string) error {
	meta := &pb.ConvertRequest{
		Value: &pb.ConvertRequest_Meta{Meta: &pb.Meta{Id: id, Type: "png", Quality: "90"}},
	}
	stream.Send(meta)

	buf := make([]byte, bufSize)
	for {
		n, err := src.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		data := &pb.ConvertRequest{
			Value: &pb.ConvertRequest_Chunk{Chunk: &pb.Chunk{Data: buf, Position: int64(n)}},
		}
		err = stream.Send(data)
		if err != nil {
			return err
		}
	}

	err := stream.CloseSend()
	if err != nil {
		return err
	}

	return nil
}

func receive(stream pb.Converter_ConvertClient, dst io.Writer) error {
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		_, err = dst.Write(resp.Data)
		if err != nil {
			return err
		}
	}

	return nil
}
