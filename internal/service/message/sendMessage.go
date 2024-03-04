package message

import (
	"context"

	model "github.com/Arturyus92/chat-server/internal/models"
)

// SendMessage - ...
func (s *serv) SendMessage(ctx context.Context, message *model.Message) error {

	err := s.messageRepository.SendMessage(ctx, message)

	if err != nil {
		return err
	}

	return nil
}
