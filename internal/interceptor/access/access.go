package access

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const accessToken = ""

// Access - ...
func (i *accessInterceptor) Access(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md := metadata.New(map[string]string{"Autorization": "Bearer " + accessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)

	err := i.client.Check(ctx, info.FullMethod)
	if err != nil {
		return nil, err
	}

	return handler(ctx, req)
}
