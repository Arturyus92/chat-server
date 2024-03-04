package model

import "time"

// Message - ...
type Message struct {
	MessageID   int64
	ChatID      int64
	UserID      int64
	TextMessage string
	CreatedAt   time.Time
}
