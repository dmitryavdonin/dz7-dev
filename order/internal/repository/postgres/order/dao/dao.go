package dao

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Id           uuid.UUID `db:"id"`
	ProductId    int       `db:"product_id"`
	ProductCount int       `db:"product_count"`
	ProductPrice float32   `json:"product_price"`
	CreatedAt    time.Time `db:"created_at"`
	ModifiedAt   time.Time `db:"modified_at"`
}

var OrderColumns = []string{
	"id",
	"product_id",
	"product_count",
	"product_price",
	"created_at",
	"modified_at",
}
