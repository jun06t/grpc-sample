package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	pb "github.com/jun06t/grpc-sample/image-proxy/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type server struct{}

func (s *server) Convert(stream pb.Converter_ConvertServer) error {
	var (
		file *os.File
		err  error
	)
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if meta := resp.GetMeta(); meta != nil {
			file, err = os.Create(fmt.Sprintf("%s.jpg", meta.Id))
			if err != nil {
				return err
			}
			defer file.Close()
		}
		if chunk := resp.GetChunk(); chunk != nil {
			file.Write(chunk.Data)
		}
	}

	err = stream.SendAndClose(&pb.ConvertResponse{"success"})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterConverterServer(s, new(server))
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
