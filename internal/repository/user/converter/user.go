package converter

import (
	model "github.com/Arturyus92/chat-server/internal/models"
	modelRepo "github.com/Arturyus92/chat-server/internal/repository/user/model"
)

// ToUserFromRepo - ...
func ToUserFromRepo(user *modelRepo.User) *model.User {
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}
}
