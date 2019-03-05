package interceptor

import (
	"context"
	"log"

	"github.com/lithammer/shortuuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func ClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		md := metadata.Pairs("x-request-id", shortuuid.New())
		ctx = metadata.NewOutgoingContext(ctx, md)

		var header, trailer metadata.MD
		opts = append(opts, grpc.Header(&header))
		opts = append(opts, grpc.Trailer(&trailer))
		err := invoker(ctx, method, req, reply, cc, opts...)
		log.Printf("reply: %+v\n", reply)
		log.Printf("header: %+v\n", header)
		log.Printf("trailer: %+v\n", trailer)
		return err
	}
}
