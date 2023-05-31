package repository

import (
	"context"
	"order/internal/domain/order"

	"github.com/google/uuid"
)

type Order interface {
	CreateOrder(ctx context.Context, order *order.Order) (err error)
	UpdateOrder(ctx context.Context, id uuid.UUID, upFunc func(oldOrder *order.Order) (*order.Order, error)) (order *order.Order, err error)
	DeleteOrderById(ctx context.Context, id uuid.UUID) (err error)
	ReadOrderById(ctx context.Context, id uuid.UUID) (order *order.Order, err error)
	ReadOrderByMsgId(ctx context.Context, msg_id uuid.UUID) (order *order.Order, err error)
}
