package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	pb "github.com/jun06t/grpc-sample/fieldmask/proto/go/user"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

const (
	address = "localhost:8080"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := newClient(conn)
	defer c.Close()

	c.Save()
	c.Get()
	c.GetWithMask()
	c.SaveWithMask()
	c.Get()
}

type Client struct {
	conn *grpc.ClientConn
	cli  pb.UserServiceClient
}

func newClient(conn *grpc.ClientConn) *Client {
	cli := pb.NewUserServiceClient(conn)
	return &Client{
		conn: conn,
		cli:  cli,
	}
}

func (c *Client) Get() error {
	req := &pb.GetRequest{
		Id: "001",
	}
	resp, err := c.cli.Get(context.Background(), req)
	if err != nil {
		return err
	}

	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
	return nil
}

func (c *Client) Save() error {
	req := &pb.UpdateRequest{
		User: &pb.User{
			Id:    "001",
			Name:  "alice",
			Age:   20,
			Email: "alice@gmail.com",
			Address: &pb.Address{
				Country: "Japan",
				State:   "Tokyo",
				City:    "Shibuya",
				Zipcode: "150-0000",
			},
		},
	}
	_, err := c.cli.Update(context.Background(), req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetWithMask() error {
	paths := []string{"name", "age", "address.city"}
	fm := fieldmaskpb.FieldMask{Paths: paths}
	req := &pb.GetRequest{
		Id:        "001",
		FieldMask: &fm,
	}
	resp, err := c.cli.Get(context.Background(), req)
	if err != nil {
		return err
	}
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
	return nil
}

func (c *Client) SaveWithMask() error {
	paths := []string{"id", "name", "email", "address.city", "address.zipcode"}
	fm := fieldmaskpb.FieldMask{Paths: paths}
	req := &pb.UpdateRequest{
		User: &pb.User{
			Id:    "001",
			Name:  "bob",
			Email: "",
			Age:   25,
			Address: &pb.Address{
				Country: "Japan",
				State:   "Tokyo",
				City:    "Ikebukuro",
				Zipcode: "170-0000",
			},
		},
		FieldMask: &fm,
	}
	_, err := c.cli.Update(context.Background(), req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Close() {
	c.conn.Close()
}
