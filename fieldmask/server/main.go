package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mennanov/fmutils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/jun06t/grpc-sample/fieldmask/proto/go/user"
)

const (
	port = ":8080"
)

type server struct {
	mcli *mongoClient
}

func (s *server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetReply, error) {
	user, err := s.mcli.GetUser(ctx, in.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to get user: %s", err))
	}
	resp := &pb.GetReply{
		User: s.toUserProto(user),
	}
	if in.FieldMask.IsValid(resp.User) {
		fmutils.Filter(resp.User, in.FieldMask.GetPaths())
	}
	return resp, nil
}

func (s *server) Update(ctx context.Context, in *pb.UpdateRequest) (*empty.Empty, error) {
	if in.FieldMask.IsValid(in.User) {
		fmutils.Filter(in.User, in.FieldMask.GetPaths())
	}
	data := s.toUserEntity(in.User, in.FieldMask.GetPaths())
	err := s.mcli.UpdateUser(ctx, data)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to update user: %s", err))
	}
	return &empty.Empty{}, nil
}

func main() {
	mcli, err := newClient(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{mcli})
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *server) toUserProto(in User) *pb.User {
	return &pb.User{
		Id:    in.ID,
		Name:  in.Name,
		Age:   int64(in.Age),
		Email: *in.Email,
		Address: &pb.Address{
			Country: in.Address.Country,
			State:   in.Address.State,
			City:    in.Address.City,
			Zipcode: in.Address.Zipcode,
		},
	}
}

func (s *server) toUserEntity(in *pb.User, paths []string) User {
	u := User{
		ID:    in.Id,
		Name:  in.Name,
		Age:   int(in.Age),
		Email: &in.Email,
		Address: Address{
			Country: in.Address.Country,
			State:   in.Address.State,
			City:    in.Address.City,
			Zipcode: in.Address.Zipcode,
		},
	}
	if len(paths) == 0 {
		return u
	}

	// set nil to omit empty
	u.Email = nil

	for i := range paths {
		if paths[i] == "email" {
			u.Email = &in.Email
		}
	}
	return u
}
