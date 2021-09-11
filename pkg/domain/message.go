package domain

import (
	"time"
)

type Message struct {
	FromUserID string
	ToChatName string
	Text       string
	Time       time.Time
}
