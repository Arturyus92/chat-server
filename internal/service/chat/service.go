package chat

import (
	"github.com/Arturyus92/chat-server/internal/repository"
	"github.com/Arturyus92/chat-server/internal/service"
	"github.com/Arturyus92/platform_common/pkg/db"
)

var _ service.ChatService = (*serv)(nil)

type serv struct {
	chatRepository     repository.ChatRepository
	chatUserRepository repository.ChatUserRepository
	logRepository      repository.LogRepository
	txManager          db.TxManager
}

// NewService - ...
func NewService(chatRepository repository.ChatRepository,
	chatUserRepository repository.ChatUserRepository,
	logRepository repository.LogRepository,
	txManager db.TxManager) *serv {
	return &serv{
		chatRepository:     chatRepository,
		chatUserRepository: chatUserRepository,
		logRepository:      logRepository,
		txManager:          txManager,
	}
}

// NewMockService - ...
func NewMockService(deps ...interface{}) service.ChatService {
	srv := serv{}

	for _, v := range deps {
		switch s := v.(type) {
		case repository.ChatRepository:
			srv.chatRepository = s
		}
	}
	return srv.chatRepository
}
