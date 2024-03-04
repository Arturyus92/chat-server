package chat

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"

	"github.com/Arturyus92/chat-server/internal/client/db"
	model "github.com/Arturyus92/chat-server/internal/models"
)

const (
	tableName = "chats"

	colChatID = "chat_id"
	colName   = "chat_title"
)

// Repo - ...
type Repo struct {
	db db.Client
}

// NewRepository - ...
func NewRepository(db db.Client) *Repo {
	return &Repo{db: db}
}

// Create - ...
func (r *Repo) Create(ctx context.Context, chat *model.Chat) (int64, error) {
	// Делаем запрос на вставку записи в таблицу chats
	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(colName).
		Values(chat.Title).
		Suffix("RETURNING " + colChatID)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return 0, err
	}
	q := db.Query{
		Name:     "chat_repository.Create",
		QueryRaw: query,
	}
	var chatID int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&chatID)
	if err != nil {
		log.Printf("failed to created chat: %v", err)
		return 0, err
	}

	return chatID, nil
}

// Delete - ...
func (r *Repo) Delete(ctx context.Context, id int64) error {
	// Делаем запрос на вставку записи в таблицу chats
	builderDelete := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{colChatID: id})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return err
	}
	q := db.Query{
		Name:     "chat_repository.Delete",
		QueryRaw: query,
	}
	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
