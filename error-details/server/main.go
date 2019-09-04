package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "github.com/jun06t/grpc-sample/error-details/proto"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	port = ":8080"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	n := rand.Intn(5) // エラーレスポンスを返すためのランダム値
	switch n {
	case 0:
		st := status.New(codes.InvalidArgument, "some error occurred")
		v := &errdetails.QuotaFailure{
			Violations: []*errdetails.QuotaFailure_Violation{
				{
					Subject:     "clientip:<ip address of client>",
					Description: "Daily Limit was exceeded for SayHello",
				},
			},
		}
		dt, _ := st.WithDetails(v)
		return nil, dt.Err()
	case 1:
		st := status.New(codes.InvalidArgument, "some error occurred")
		v := &errdetails.BadRequest{
			FieldViolations: []*errdetails.BadRequest_FieldViolation{
				{
					Field:       "username",
					Description: "should not empty",
				},
			},
		}
		dt, _ := st.WithDetails(v)
		return nil, dt.Err()
	case 2:
		st := status.New(codes.InvalidArgument, "some error occurred")
		dt, _ := st.WithDetails(&pb.ErrorDetail{Code: pb.ErrorCode_EXPIRED_RECEIPT})
		return nil, dt.Err()
	case 3:
		st := status.New(codes.InvalidArgument, "some error occurred")
		dt, _ := st.WithDetails(&pb.ErrorDetail{Code: pb.ErrorCode_INVALID_COUNTRY})
		return nil, dt.Err()
	}

	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
