package access

import (
	"context"

	access "github.com/Arturyus92/chat-server/pkg/access_v1"
)

type accessClient struct {
	client access.AccessV1Client
}

// NewAccessClient ...
func NewAccessClient(client access.AccessV1Client) *accessClient {
	return &accessClient{
		client: client,
	}
}

// Check ...
func (a *accessClient) Check(ctx context.Context, address string) error {
	_, err := a.client.Check(
		ctx, &access.CheckRequest{
			EndpointAddress: address,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
