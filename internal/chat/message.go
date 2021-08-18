package chat

import (
	"time"
)

type Message struct {
	FromID string
	ToID   string
	Text   string
	Time   time.Time
}
