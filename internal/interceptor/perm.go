package interceptor

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Client struct {
	Client rpc.AuthClient
}

// PermissionInterceptor - ...
func (c *Client) PermissionInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata is false")
	}

	err := c.Client.Check(metadata.NewOutgoingContext(ctx, md), info.FullMethod)
	if err != nil {
		return nil, err
	}

	return handler(ctx, req)
}
