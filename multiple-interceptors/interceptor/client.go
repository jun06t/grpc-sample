package interceptor

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func ClientInterceptor(prefix string) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		fmt.Println(prefix, "before invoker")
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Println(prefix, "after invoker")
		return err
	}
}
