package chat

import (
	"github.com/Arturyus92/chat-server/internal/service"
	desc "github.com/Arturyus92/chat-server/pkg/chat_v1"
)

// Implementation - ...
type Implementation struct {
	desc.UnimplementedChatV1Server
	messageService service.MessageService
	chatService    service.ChatService
}

// NewImplementation - ...
func NewImplementation(messageService service.MessageService, chatService service.ChatService) *Implementation {
	return &Implementation{
		messageService: messageService,
		chatService:    chatService,
	}
}
