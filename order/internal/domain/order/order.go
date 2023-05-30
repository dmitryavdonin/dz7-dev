package order

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	id            uuid.UUID
	product_id    int
	produtc_count int
	product_price float32
	createdAt     time.Time
	modifiedAt    time.Time
}

func NewOrder(
	product_id int,
	produtc_count int,
	product_price float32,
) (*Order, error) {

	return &Order{
		id:            uuid.New(),
		product_id:    product_id,
		produtc_count: produtc_count,
		product_price: product_price,
		createdAt:     time.Now(),
		modifiedAt:    time.Now(),
	}, nil
}

func NewOrderWithId(
	id uuid.UUID,
	product_id int,
	produtc_count int,
	product_price float32,
	createdAt time.Time,
	modifiedAt time.Time,
) *Order {
	return &Order{
		id:            id,
		product_id:    product_id,
		produtc_count: produtc_count,
		product_price: product_price,
		createdAt:     createdAt,
		modifiedAt:    modifiedAt,
	}
}

func (u Order) Id() uuid.UUID {
	return u.id
}

func (u Order) ProductId() int {
	return u.product_id
}

func (p Order) ProductCount() int {
	return p.produtc_count
}

func (p Order) ProductPrice() float32 {
	return p.product_price
}

func (u Order) CreatedAt() time.Time {
	return u.createdAt
}

func (u Order) ModifiedAt() time.Time {
	return u.modifiedAt
}
