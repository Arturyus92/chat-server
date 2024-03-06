package converter

import (
	model "github.com/Arturyus92/chat-server/internal/models"
	desc "github.com/Arturyus92/chat-server/pkg/chat_v1"
)

// ToChatCreateFromDesc - ...
func ToChatCreateFromDesc(chat *desc.ChatInfo) *model.Chat {
	return &model.Chat{
		Users: chat.Usernames,
		Title: chat.Title,
	}
}
