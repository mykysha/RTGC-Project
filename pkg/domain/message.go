package domain

import (
	"time"
)

type Message struct {
	FromUserName string
	ToRoomName   string
	ToID         []string
	Text         string
	Time         time.Time
}
