package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"

	"github.com/stretchr/testify/require"

	"github.com/Arturyus92/chat-server/internal/repository"
	repoMocks "github.com/Arturyus92/chat-server/internal/repository/mocks"
	"github.com/Arturyus92/chat-server/internal/service/chat"
)

func TestDelete(t *testing.T) {
	t.Parallel()
	type chatRepositoryMockFunc func(mc *minimock.Controller) repository.ChatRepository

	type args struct {
		ctx context.Context
		id  int64
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id      = gofakeit.Int64()
		repoErr = fmt.Errorf("repo error")
	)

	tests := []struct {
		name               string
		args               args
		want               int64
		err                error
		chatRepositoryMock chatRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				id:  id,
			},
			err: nil,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				id:  id,
			},
			err: repoErr,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(repoErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatRepoMock := tt.chatRepositoryMock(mc)
			service := chat.NewMockService(chatRepoMock)

			err := service.Delete(tt.args.ctx, tt.args.id)
			require.Equal(t, tt.err, err)
		})
	}
}
