package dao

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"created_at"`
}

var OrderColumns = []string{
	"id",
	"created_at",
}
