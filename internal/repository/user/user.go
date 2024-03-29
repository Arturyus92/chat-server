package user

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"

	model "github.com/Arturyus92/chat-server/internal/models"
	"github.com/Arturyus92/platform_common/pkg/db"
)

const (
	tableName = "users"

	colUserID = "user_id"
	colName   = "name"
)

// Repo - ...
type Repo struct {
	db db.Client
}

// NewRepository - ...
func NewRepository(db db.Client) *Repo {
	return &Repo{db: db}
}

// Get - ...
func (r *Repo) Get(ctx context.Context, user *model.User) (int64, error) {
	// Делаем запрос на получение записи из таблицы user
	builderSelectOne := sq.Select(colUserID).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{colName: user.Name}).
		Limit(1)

	query, args, err := builderSelectOne.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return 0, err
	}
	q := db.Query{
		Name:     "user_repository.Get",
		QueryRaw: query,
	}

	var userID int64
	err = r.db.DB().ScanOneContext(ctx, &userID, q, args...)
	if err != nil {
		log.Printf("failed to ScanOneContext: %v", err)
		return 0, err
	}

	return userID, nil
}
