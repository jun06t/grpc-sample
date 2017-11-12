package main

import (
	"log"
	"net"
	"time"

	pb "github.com/jun06t/grpc-sample/server-streaming/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type server struct{}

func (s *server) GetNewFeed(in *pb.Empty, stream pb.Feeder_GetNewFeedServer) error {
	feed := []string{"article1", "article2", "article3"}

	for _, v := range feed {
		err := stream.Send(&pb.FeedResponse{Message: v})
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterFeederServer(s, new(server))
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
