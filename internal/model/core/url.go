package core

import (
	"time"
)

type URL struct {
	Hash       string
	Value      string
	CreatedAt  time.Time
	ExpirestAt time.Time
}
