package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

// AccessInterceptor ...
type AccessInterceptor interface {
	Access(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (
		resp interface{}, err error,
	)
}
