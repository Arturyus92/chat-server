package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"

	"github.com/stretchr/testify/require"

	chatAPI "github.com/Arturyus92/chat-server/internal/api"
	model "github.com/Arturyus92/chat-server/internal/models"
	"github.com/Arturyus92/chat-server/internal/service"
	serviceMocks "github.com/Arturyus92/chat-server/internal/service/mocks"
	desc "github.com/Arturyus92/chat-server/pkg/chat_v1"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService
	type messageServiceMockFunc func(mc *minimock.Controller) service.MessageService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id    = gofakeit.Int64()
		title = gofakeit.Animal()
		users = []int64{gofakeit.Int64(), gofakeit.Int64(), gofakeit.Int64()}

		serviceErr = fmt.Errorf("service error")

		req = &desc.CreateRequest{
			ChatInfo: &desc.ChatInfo{
				UserId: users,
				Title:  title,
			},
		}

		chatInfo = &model.Chat{
			Title: title,
			Users: users,
		}

		res = &desc.CreateResponse{
			Id: id,
		}
	)

	tests := []struct {
		name               string
		args               args
		want               *desc.CreateResponse
		err                error
		chatServiceMock    chatServiceMockFunc
		messageServiceMock messageServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.CreateMock.Expect(ctx, chatInfo).Return(id, nil)
				return mock
			},
			messageServiceMock: func(mc *minimock.Controller) service.MessageService {
				mock := serviceMocks.NewMessageServiceMock(mc)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.CreateMock.Expect(ctx, chatInfo).Return(0, serviceErr)
				return mock
			},
			messageServiceMock: func(mc *minimock.Controller) service.MessageService {
				mock := serviceMocks.NewMessageServiceMock(mc)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatServiceMock := tt.chatServiceMock(mc)
			messageServiceMock := tt.messageServiceMock(mc)
			api := chatAPI.NewImplementation(messageServiceMock, chatServiceMock)

			res, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}
