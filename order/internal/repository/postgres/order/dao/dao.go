package dao

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Id           uuid.UUID `db:"id"`
	MsgId        uuid.UUID `db:"msg_id"`
	ProductId    int       `db:"product_id"`
	ProductCount int       `db:"product_count"`
	ProductPrice float32   `db:"product_price"`
	Version      int       `db:"version"`
	CreatedAt    time.Time `db:"created_at"`
	ModifiedAt   time.Time `db:"modified_at"`
}

var OrderColumns = []string{
	"id",
	"msg_id",
	"product_id",
	"product_count",
	"product_price",
	"version",
	"created_at",
	"modified_at",
}
