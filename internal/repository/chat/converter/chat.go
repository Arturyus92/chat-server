package converter

import (
	model "github.com/Arturyus92/chat-server/internal/models"

	modelRepo "github.com/Arturyus92/chat-server/internal/repository/chat/model"
)

// ToChatFromRepo - ...
func ToChatFromRepo(chat *modelRepo.Chat) *model.Chat {
	return &model.Chat{
		ID:    chat.ID,
		Title: chat.Title,
	}
}
