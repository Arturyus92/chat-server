package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"

	"github.com/stretchr/testify/require"

	model "github.com/Arturyus92/chat-server/internal/models"
	"github.com/Arturyus92/chat-server/internal/repository"
	repoMocks "github.com/Arturyus92/chat-server/internal/repository/mocks"
	"github.com/Arturyus92/chat-server/internal/service/message"
)

func TestSendMessage(t *testing.T) {
	t.Parallel()
	type messageRepositoryMockFunc func(mc *minimock.Controller) repository.MessageRepository

	type args struct {
		ctx context.Context
		req *model.Message
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id     = gofakeit.Int64()
		from   = gofakeit.Int64()
		chatID = gofakeit.Int64()
		text   = gofakeit.Animal()

		repoErr = fmt.Errorf("repo error")

		req = &model.Message{
			MessageID:   id,
			ChatID:      chatID,
			UserID:      from,
			TextMessage: text,
		}
	)

	tests := []struct {
		name                  string
		args                  args
		want                  int64
		err                   error
		messageRepositoryMock messageRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: id,
			err:  nil,
			messageRepositoryMock: func(mc *minimock.Controller) repository.MessageRepository {
				mock := repoMocks.NewMessageRepositoryMock(mc)
				mock.CreateMessageMock.Expect(ctx, req).Return(nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: 0,
			err:  repoErr,
			messageRepositoryMock: func(mc *minimock.Controller) repository.MessageRepository {
				mock := repoMocks.NewMessageRepositoryMock(mc)
				mock.CreateMessageMock.Expect(ctx, req).Return(repoErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			messageRepoMock := tt.messageRepositoryMock(mc)
			service := message.NewMockService(messageRepoMock)

			err := service.SendMessage(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}
