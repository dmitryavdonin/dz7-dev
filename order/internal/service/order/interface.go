package user

import (
	"context"
	"order/internal/domain/order"

	"github.com/google/uuid"
)

type Service interface {
	CreateOrder(ctx context.Context, order *order.Order) (err error)
	UpdateOrder(ctx context.Context, id uuid.UUID, upFn func(oldOrder *order.Order) (*order.Order, error)) (*order.Order, error)
	DeleteOrderById(ctx context.Context, id uuid.UUID) (err error)
	ReadOrderById(ctx context.Context, id uuid.UUID) (order *order.Order, err error)
	ReadOrderByMsgId(ctx context.Context, msg_id uuid.UUID) (order *order.Order, err error)
}
