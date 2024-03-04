package chat

import (
	"context"
	"fmt"

	model "github.com/Arturyus92/chat-server/internal/models"
)

// Delete ...
func (s *serv) Delete(ctx context.Context, id int64) error {

	err := s.txManager.ReadCommitted(
		ctx, func(ctx context.Context) error {
			errTx := s.chatRepository.Delete(ctx, id)
			if errTx != nil {
				return errTx
			}

			errTx = s.logRepository.CreateLog(ctx, &model.Log{
				Info: fmt.Sprintf("Chat deleteted: %d", id),
			})
			if errTx != nil {
				return errTx
			}

			return nil
		},
	)
	if err != nil {
		return err
	}

	return nil
}
