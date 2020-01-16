package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/jun06t/grpc-sample/client-side-lb/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
)

var (
	endpoint = "localhost:8080"
)

func init() {
	ep := os.Getenv("ENDPOINT")
	if ep != "" {
		endpoint = ep
	}
}

func main() {
	fmt.Println("Endpoint: ", endpoint)
	resolver.SetDefaultScheme("dns")
	conn, err := grpc.Dial(endpoint,
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	req := &pb.HelloRequest{
		Name: "alice",
	}
	for {
		resp, err := c.SayHello(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Machine: %s, Reply: %s\n", resp.MachineId, resp.Message)
		time.Sleep(1 * time.Second)
	}
}
