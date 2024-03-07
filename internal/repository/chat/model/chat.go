package model

// Chat - ...
type Chat struct {
	ID    int64  `db:"chat_id"`
	Title string `db:"chat_title"`
	Users []string
}
