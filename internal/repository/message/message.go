package message

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"

	"github.com/Arturyus92/chat-server/internal/client/db"
	model "github.com/Arturyus92/chat-server/internal/models"
)

const (
	tableName      = "messages"
	colMessageID   = "message_id"
	colChatID      = "chat_id"
	colUserID      = "user_id"
	colTextMessage = "text_message"
)

// Repo - ...
type Repo struct {
	db db.Client
}

// NewRepository - ...
func NewRepository(db db.Client) *Repo {
	return &Repo{db: db}
}

// CreateMessage - ...
func (r *Repo) CreateMessage(ctx context.Context, message *model.Message) error {
	// Делаем запрос на вставку записи в таблицу messages
	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(colChatID, colUserID, colTextMessage).
		Values(message.ChatID, message.UserID, message.TextMessage).
		Suffix("RETURNING " + colUserID)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return err
	}
	q := db.Query{
		Name:     "message_repository.Create",
		QueryRaw: query,
	}

	var messageID int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&messageID)
	if err != nil {
		log.Printf("failed to created message: %v", err)
		return err
	}

	return nil
}
