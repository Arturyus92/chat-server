package model

import "time"

// Message - ...
type Message struct {
	MessageID   int64     `db:"message_id"`
	ChatID      int64     `db:"chat_id"`
	UserID      int64     `db:"user_id"`
	TextMessage string    `db:"text_message"`
	CreatedAt   time.Time `db:"created_at"`
}
