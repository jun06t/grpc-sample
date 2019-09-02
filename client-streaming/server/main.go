package main

import (
	"io"
	"log"
	"net"
	"os"

	pb "github.com/jun06t/grpc-sample/client-streaming/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type server struct{}

func (s *server) Upload(stream pb.Uploader_UploadServer) error {
	file, err := os.Create("supercar.jpg")
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

	err = stream.SendAndClose(&pb.UploadResponse{Status: "success"})
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
	pb.RegisterUploaderServer(s, new(server))
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
