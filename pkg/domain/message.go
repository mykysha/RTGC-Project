package domain

import (
	"time"
)

type Message struct {
	FromUserID string
	ToRoomName string
	ToID       []string
	Text       string
	Time       time.Time
}
