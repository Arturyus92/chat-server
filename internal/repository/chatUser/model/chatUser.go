package model

// ChatUser - ...
type ChatUser struct {
	ChatID int64 `db:"chat_id"`
	UserID int64 `db:"user_id"`
}
