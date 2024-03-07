package user

import (
	"context"

	model "github.com/Arturyus92/chat-server/internal/models"
)

// Get - ...
func (s *serv) Get(ctx context.Context, user *model.User) (int64, error) {
	id, err := s.userRepository.Get(ctx, user)
	if err != nil {
		return 0, err
	}

	if err != nil {
		return 0, err
	}

	return id, nil
}
