package converter

import (
	model "github.com/Arturyus92/chat-server/internal/models"
	modelRepo "github.com/Arturyus92/chat-server/internal/repository/chatUser/model"
)

// ToChatUserFromRepo - ...
func ToChatUserFromRepo(chatUser *modelRepo.ChatUser) *model.ChatUser {
	return &model.ChatUser{
		ChatID: chatUser.ChatID,
		UserID: chatUser.UserID,
	}
}
