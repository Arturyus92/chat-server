package app

import (
	"context"
	"log"

	chat "github.com/Arturyus92/chat-server/internal/api"
	rpc "github.com/Arturyus92/chat-server/internal/client"
	accessClient "github.com/Arturyus92/chat-server/internal/client/rpc/access"
	"github.com/Arturyus92/chat-server/internal/config"
	"github.com/Arturyus92/chat-server/internal/config/env"
	"github.com/Arturyus92/chat-server/internal/interceptor"
	accessInterceptor "github.com/Arturyus92/chat-server/internal/interceptor/access"
	"github.com/Arturyus92/chat-server/internal/repository"
	chatRepository "github.com/Arturyus92/chat-server/internal/repository/chat"
	chatUserRepository "github.com/Arturyus92/chat-server/internal/repository/chat_user"
	logRepository "github.com/Arturyus92/chat-server/internal/repository/log"
	messageRepository "github.com/Arturyus92/chat-server/internal/repository/message"
	userRepository "github.com/Arturyus92/chat-server/internal/repository/user"
	"github.com/Arturyus92/chat-server/internal/service"
	chatService "github.com/Arturyus92/chat-server/internal/service/chat"
	messageService "github.com/Arturyus92/chat-server/internal/service/message"
	access "github.com/Arturyus92/chat-server/pkg/access_v1"
	"github.com/Arturyus92/platform_common/pkg/closer"
	"github.com/Arturyus92/platform_common/pkg/db"
	"github.com/Arturyus92/platform_common/pkg/db/pg"
	"github.com/Arturyus92/platform_common/pkg/db/transaction"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const authServiceAddress = "auth_service:50052"

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient           db.Client
	txManager          db.TxManager
	userRepository     repository.UserRepository
	chatRepository     repository.ChatRepository
	chatUserRepository repository.ChatUserRepository
	messageRepository  repository.MessageRepository
	logRepository      repository.LogRepository

	chatService    service.ChatService
	messageService service.MessageService

	chatImpl *chat.Implementation

	accessClient      rpc.AccessClient
	accessInterceptor interceptor.AccessInterceptor
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// PGConfig - ...
func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

// GRPCConfig - ...
func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

// DBClient - ...
func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

// TxManager - ...
func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

// ChatRepository - ...
func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepository.NewRepository(s.DBClient(ctx))
	}

	return s.chatRepository
}

// MessageRepository - ...
func (s *serviceProvider) MessageRepository(ctx context.Context) repository.MessageRepository {
	if s.messageRepository == nil {
		s.messageRepository = messageRepository.NewRepository(s.DBClient(ctx))
	}

	return s.messageRepository
}

// UserRepository - ...
func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.DBClient(ctx))
	}

	return s.userRepository
}

// ChatUserRepository - ...
func (s *serviceProvider) ChatUserRepository(ctx context.Context) repository.ChatUserRepository {
	if s.chatUserRepository == nil {
		s.chatUserRepository = chatUserRepository.NewRepository(s.DBClient(ctx))
	}

	return s.chatUserRepository
}

// LogRepository - ...
func (s *serviceProvider) LogRepository(ctx context.Context) repository.LogRepository {
	if s.logRepository == nil {
		s.logRepository = logRepository.NewRepository(s.DBClient(ctx))
	}

	return s.logRepository
}

// ChatService - ...
func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewService(
			s.ChatRepository(ctx),
			s.ChatUserRepository(ctx),
			s.LogRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.chatService
}

// MessageService - ...
func (s *serviceProvider) MessageService(ctx context.Context) service.MessageService {
	if s.messageService == nil {
		s.messageService = messageService.NewService(
			s.MessageRepository(ctx),
		)
	}

	return s.messageService
}

// ChatImpl - ...
func (s *serviceProvider) ChatImpl(ctx context.Context) *chat.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chat.NewImplementation(s.MessageService(ctx), s.ChatService(ctx))
	}

	return s.chatImpl
}

func (s *serviceProvider) AccessInterceptor(ctx context.Context) interceptor.AccessInterceptor {
	if s.accessInterceptor == nil {
		s.accessInterceptor = accessInterceptor.NewAccessInterceptor(s.AccessClient(ctx))
	}

	return s.accessInterceptor
}

func (s *serviceProvider) AccessClient(ctx context.Context) rpc.AccessClient {
	if s.accessClient == nil {
		//conn, err := grpc.DialContext(ctx, authServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
		conn, err := grpc.Dial(authServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("init AccessClient error")
		}

		closer.Add(conn.Close)

		s.accessClient = accessClient.NewAccessClient(access.NewAccessV1Client(conn))
	}

	return s.accessClient
}
