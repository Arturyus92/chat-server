package service

import (
	"context"

	model "github.com/Arturyus92/chat-server/internal/models"
)

// ChatService - ...
type ChatService interface {
	Create(ctx context.Context, chat *model.Chat) (int64, error)
	Delete(ctx context.Context, id int64) error
}

// MessageService - ...
type MessageService interface {
	SendMessage(ctx context.Context, message *model.Message) error
}
