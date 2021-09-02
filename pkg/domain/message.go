package domain

import (
	"time"
)

type Message struct {
	FromUserID string
	ToGroupID  string
	Text       string
	Time       time.Time
}
