package order

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	id            uuid.UUID
	msg_id        uuid.UUID
	product_id    int
	product_count int
	product_price float32
	version       int
	createdAt     time.Time
	modifiedAt    time.Time
}

func NewOrder(
	msg_id uuid.UUID,
	product_id int,
	produtc_count int,
	product_price float32,
) (*Order, error) {

	return &Order{
		id:            uuid.New(),
		msg_id:        msg_id,
		product_id:    product_id,
		product_count: produtc_count,
		product_price: product_price,
		version:       0,
		createdAt:     time.Now(),
		modifiedAt:    time.Now(),
	}, nil
}

func NewOrderWithId(
	id uuid.UUID,
	msg_id uuid.UUID,
	product_id int,
	produtc_count int,
	product_price float32,
	version int,
	createdAt time.Time,
	modifiedAt time.Time,
) *Order {
	return &Order{
		id:            id,
		msg_id:        msg_id,
		product_id:    product_id,
		product_count: produtc_count,
		product_price: product_price,
		version:       version,
		createdAt:     createdAt,
		modifiedAt:    modifiedAt,
	}
}

func (u Order) Id() uuid.UUID {
	return u.id
}

func (u Order) MsgId() uuid.UUID {
	return u.msg_id
}

func (u Order) Version() int {
	return u.version
}

func (u Order) ProductId() int {
	return u.product_id
}

func (p Order) ProductCount() int {
	return p.product_count
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
