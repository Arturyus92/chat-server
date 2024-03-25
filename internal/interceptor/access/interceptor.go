package access

import rpc "github.com/Arturyus92/chat-server/internal/client"

type accessInterceptor struct {
	client rpc.AccessClient
}

// NewAccessInterceptor ...
func NewAccessInterceptor(rpcClient rpc.AccessClient) *accessInterceptor {
	return &accessInterceptor{
		client: rpcClient,
	}
}
