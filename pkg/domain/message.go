package domain

import (
	"time"
)

type Message struct {
	FromUserID string
	ToChatID   string
	Text       string
	Time       time.Time
}
