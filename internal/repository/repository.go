package repository

import (
	"context"

	model "github.com/Arturyus92/chat-server/internal/models"
)

// UserRepository - ...
type UserRepository interface {
	Get(ctx context.Context, user *model.User) (int64, error)
}

// ChatUserRepository - ...
type ChatUserRepository interface {
	CreateChat(ctx context.Context, chatUser *model.ChatUser) error
}

// ChatRepository - ...
type ChatRepository interface {
	Create(ctx context.Context, chat *model.Chat) (int64, error)
	Delete(ctx context.Context, id int64) error
}

// MessageRepository - ...
type MessageRepository interface {
	SendMessage(ctx context.Context, message *model.Message) error
}

// LogRepository - ...
type LogRepository interface {
	CreateLog(ctx context.Context, log *model.Log) error
}
