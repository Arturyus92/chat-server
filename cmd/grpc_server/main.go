package main

import (
	"context"
	"flag"
	"log"
	"net"

	sq "github.com/Masterminds/squirrel"
	"github.com/brianvoe/gofakeit"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Arturyus92/chat-server/internal/config"
	"github.com/Arturyus92/chat-server/internal/config/env"
	desc "github.com/Arturyus92/chat-server/pkg/chat_v1"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "prod.env", "path to config file")
}

type server struct {
	desc.UnimplementedChatV1Server
	pool *pgxpool.Pool
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	// Делаем запрос на вставку записи в таблицу chats
	builderInsert := sq.Insert("chats").
		PlaceholderFormat(sq.Dollar).
		Columns("chat_title").
		Values(gofakeit.BeerName()). //надо proto-файл менять, пока так
		Suffix("RETURNING chat_id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
	}

	var chatID int64
	err = s.pool.QueryRow(ctx, query, args...).Scan(&chatID)
	if err != nil {
		log.Printf("failed to created chat: %v", err)
	}

	log.Printf("Chat created: %+v", req.String())

	userNames := make([]string, len(req.Usernames))
	userNames = append(userNames, req.Usernames...)

	log.Printf("Users added: %+v", userNames)
	/*

		Место для вашей рекламы

	*/

	return &desc.CreateResponse{
		Id: chatID,
	}, nil

}

func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	// Достаем имя юзер из таблицы users
	builderSelectUserId := sq.Select("user_id").From("users").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"name": req.GetFrom()}).
		Limit(1)

	query, args, err := builderSelectUserId.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
	}

	var id int64

	err = s.pool.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		log.Printf("failed to selected user: %v", err)
	}

	// Делаем запрос на вставку записи в таблицу messages
	builderInsert := sq.Insert("messages").
		PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "user_id", "text_message ").
		Values(1, id, req.GetText()). //chat_id is hardcode, but this is temporary
		Suffix("RETURNING user_id")

	query, args, err = builderInsert.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
	}

	_, err = s.pool.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("failed to send message: %v", err)
	}

	log.Printf("Send message: %+v", req.GetText())

	return &emptypb.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	builderDelete := sq.Delete("chats").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"chat_id": req.Id})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
	}

	_, err = s.pool.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("failed to deleted chat: %v", err)
	}
	log.Printf("Chat deleted: %+v", req.String())

	return &emptypb.Empty{}, nil
}

func main() {
	flag.Parse()
	ctx := context.Background()

	// Считываем переменные окружения
	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Создаем пул соединений с базой данных
	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{pool: pool})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}