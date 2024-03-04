package converter

import (
	model "github.com/Arturyus92/chat-server/internal/models"

	modelRepo "github.com/Arturyus92/chat-server/internal/repository/message/model"
)

// ToMessageFromRepo - ...
func ToMessageFromRepo(message *modelRepo.Message) *model.Message {
	return &model.Message{
		MessageID:   message.MessageID,
		ChatID:      message.ChatID,
		UserID:      message.UserID,
		TextMessage: message.TextMessage,
		CreatedAt:   message.CreatedAt,
	}
}
