package domain

import (
	"time"
)

type Message struct {
	FromUserID string
	ToRoomName string
	Text       string
	Time       time.Time
}
