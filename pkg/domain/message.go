package domain

import (
	"time"
)

type Message struct {
	FromUser   string
	ToRoomName string
	ToID       []string
	Text       string
	Time       time.Time
}
