package converter

import (
	model "github.com/Arturyus92/chat-server/internal/models"
	desc "github.com/Arturyus92/chat-server/pkg/chat_v1"
)

// ToSendMessageFromDesc - ...
func ToSendMessageFromDesc(message *desc.MessageInfo) *model.Message {

	return &model.Message{
		UserID:      message.From,
		ChatID:      message.ChatId,
		TextMessage: message.Text,
	}
}
