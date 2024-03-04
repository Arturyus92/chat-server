package user

import (
	"github.com/Arturyus92/chat-server/internal/repository"
	"github.com/Arturyus92/chat-server/internal/service"
)

var _ service.UserService = (*serv)(nil)

type serv struct {
	userRepository repository.UserRepository
}

// NewService - ...
func NewService(userRepository repository.UserRepository) *serv {
	return &serv{
		userRepository: userRepository,
	}
}
