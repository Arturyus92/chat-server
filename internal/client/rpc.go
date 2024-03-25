package rpc

import "context"

//AccessClient-...
type AccessClient interface {
	Check(ctx context.Context, address string) error
}
