package message

import (
	"github.com/Arturyus92/chat-server/internal/repository"
	"github.com/Arturyus92/chat-server/internal/service"
)

var _ service.MessageService = (*serv)(nil)

type serv struct {
	messageRepository repository.MessageRepository
}

// NewService - ...
func NewService(messageRepository repository.MessageRepository) *serv {
	return &serv{
		messageRepository: messageRepository,
	}
}

// NewMockService - ...
func NewMockService(deps ...interface{}) service.MessageService {
	srv := serv{}

	for _, v := range deps {
		switch s := v.(type) {
		case repository.MessageRepository:
			srv.messageRepository = s
		}
	}

	return &srv
}
