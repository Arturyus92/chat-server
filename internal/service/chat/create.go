package chat

import (
	"context"
	"fmt"

	model "github.com/Arturyus92/chat-server/internal/models"
)

// Create ...
func (s *serv) Create(ctx context.Context, chat *model.Chat) (int64, error) {
	var chatID int64

	err := s.txManager.ReadCommitted(
		ctx, func(ctx context.Context) error {
			var errTx error
			chatID, errTx = s.chatRepository.Create(ctx, chat)
			if errTx != nil {
				return errTx
			}

			//получаем userID по имени
			for _, user := range chat.Users {
				userID, err := s.userRepository.Get(ctx, &model.User{
					Name: user,
				})
				if err != nil {
					return err
				}

				//в таблице chatUsers создаем отношение chat-user
				err = s.chatUserRepository.CreateChat(ctx, &model.ChatUser{
					ChatID: chatID,
					UserID: userID,
				})
				if err != nil {
					return err
				}
			}

			errTx = s.logRepository.CreateLog(ctx, &model.Log{
				Text: fmt.Sprintf("Chat created: %d", chatID),
			})
			if errTx != nil {
				return errTx
			}

			return nil
		},
	)
	if err != nil {
		return 0, err
	}

	return chatID, nil
}
