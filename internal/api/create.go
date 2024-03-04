package chat

import (
	"context"

	"github.com/Arturyus92/chat-server/internal/converter"
	desc "github.com/Arturyus92/chat-server/pkg/chat_v1"
)

// Create - ...
func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.chatService.Create(ctx, converter.ToChatCreateFromDesc(req.GetChatInfo()))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
