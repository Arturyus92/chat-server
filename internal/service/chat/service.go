package chat

import (
	"github.com/Arturyus92/chat-server/internal/client/db"
	"github.com/Arturyus92/chat-server/internal/repository"
	"github.com/Arturyus92/chat-server/internal/service"
)

var _ service.ChatService = (*serv)(nil)

type serv struct {
	chatRepository     repository.ChatRepository
	chatUserRepository repository.ChatUserRepository
	userRepository     repository.UserRepository
	logRepository      repository.LogRepository
	txManager          db.TxManager
}

// NewService - ...
func NewService(chatRepository repository.ChatRepository,
	chatUserRepository repository.ChatUserRepository,
	userRepository repository.UserRepository,
	logRepository repository.LogRepository,
	txManager db.TxManager) *serv {
	return &serv{
		chatRepository:     chatRepository,
		chatUserRepository: chatUserRepository,
		userRepository:     userRepository,
		logRepository:      logRepository,
		txManager:          txManager,
	}
}
