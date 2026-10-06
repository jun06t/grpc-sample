package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/jun06t/grpc-sample/client-side-lb/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver/dns"
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
	conn, err := getConn()
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

func getConn() (*grpc.ClientConn, error) {
	// The dns resolver re-resolves when a backend connection is lost.
	// Lower the rate limit so new backends are picked up quickly in this sample.
	dns.SetMinResolutionInterval(5 * time.Second)
	conn, err := grpc.NewClient("dns:///"+endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
	)
	return conn, err
}
