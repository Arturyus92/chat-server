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
	"github.com/Arturyus92/chat-server/internal/service/chat"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type chatRepositoryMockFunc func(mc *minimock.Controller) repository.ChatRepository

	type args struct {
		ctx context.Context
		req *model.Chat
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id    = gofakeit.Int64()
		title = gofakeit.Animal()
		users = []int64{gofakeit.Int64(), gofakeit.Int64(), gofakeit.Int64()}

		repoErr = fmt.Errorf("repo error")

		req = &model.Chat{
			Title: title,
			Users: users,
		}
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
				req: req,
			},
			want: id,
			err:  nil,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, req).Return(id, nil)
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
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, req).Return(0, repoErr)
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

			res, err := service.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}
