package chat

import (
	"context"

	"github.com/Arturyus92/chat-server/internal/converter"
	desc "github.com/Arturyus92/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// SendMessage - ...
func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	err := i.messageService.SendMessage(ctx, converter.ToSendMessageFromDesc(req.GetMessageInfo()))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
