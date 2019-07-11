package interceptor

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func ServerInterceptor(prefix string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		fmt.Println(prefix, "before handler")
		// call handler
		resp, err := handler(ctx, req)
		fmt.Println(prefix, "after handler")

		return resp, err
	}
}
