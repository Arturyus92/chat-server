package chatuser

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"

	"github.com/Arturyus92/chat-server/internal/client/db"
	model "github.com/Arturyus92/chat-server/internal/models"
)

const (
	tableName = "users_chats"

	colUserID = "user_id"
	colChatID = "chat_id"
)

// Repo - ...
type Repo struct {
	db db.Client
}

// NewRepository - ...
func NewRepository(db db.Client) *Repo {
	return &Repo{db: db}
}

// CreateChat - ...
func (r *Repo) CreateChat(ctx context.Context, chatUser *model.ChatUser) error {
	// Делаем запрос на вставку записи в таблицу users_chats
	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(colUserID, colChatID).
		Values(chatUser.UserID, chatUser.ChatID)

	log.Printf("User: %d, Chat: %d", chatUser.UserID, chatUser.ChatID)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return err
	}

	q := db.Query{
		Name:     "users_chats_repository.Create",
		QueryRaw: query,
	}

	var userID, chatID int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&userID, &chatID)

	if err != nil {
		log.Printf("failed to created chats_users: %v", err)
		return err
	}

	return nil
}
