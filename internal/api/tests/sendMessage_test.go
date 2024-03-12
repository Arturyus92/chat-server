package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/stretchr/testify/require"

	chatAPI "github.com/Arturyus92/chat-server/internal/api"
	model "github.com/Arturyus92/chat-server/internal/models"
	"github.com/Arturyus92/chat-server/internal/service"
	serviceMocks "github.com/Arturyus92/chat-server/internal/service/mocks"
	desc "github.com/Arturyus92/chat-server/pkg/chat_v1"
)

func TestSendMessage(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService
	type messageServiceMockFunc func(mc *minimock.Controller) service.MessageService

	type args struct {
		ctx context.Context
		req *desc.SendMessageRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		from   = gofakeit.Int64()
		chatID = gofakeit.Int64()
		text   = gofakeit.Animal()

		serviceErr = fmt.Errorf("service error")
		req        = &desc.SendMessageRequest{
			MessageInfo: &desc.MessageInfo{
				From:      from,
				Text:      text,
				Timestamp: timestamppb.New(time.Now()),
				ChatId:    chatID,
			},
		}

		messageInfo = &model.Message{
			ChatID:      chatID,
			UserID:      from,
			TextMessage: text,
		}

		res = &emptypb.Empty{}
	)

	tests := []struct {
		name               string
		args               args
		want               *emptypb.Empty
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
				return mock
			},
			messageServiceMock: func(mc *minimock.Controller) service.MessageService {
				mock := serviceMocks.NewMessageServiceMock(mc)
				mock.SendMessageMock.Expect(ctx, messageInfo).Return(nil)
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
				return mock
			},
			messageServiceMock: func(mc *minimock.Controller) service.MessageService {
				mock := serviceMocks.NewMessageServiceMock(mc)
				mock.SendMessageMock.Expect(ctx, messageInfo).Return(serviceErr)
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

			res, err := api.SendMessage(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}
