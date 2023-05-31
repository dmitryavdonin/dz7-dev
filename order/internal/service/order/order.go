package user

import (
	"context"
	"errors"
	"order/internal/domain/order"

	"github.com/google/uuid"
)

var (
	ErrOrderNotFound = errors.New("order not found")
)

func (s *service) ReadOrderById(ctx context.Context, id uuid.UUID) (order *order.Order, err error) {
	return s.repository.ReadOrderById(ctx, id)
}

func (s *service) ReadOrderByMsgId(ctx context.Context, msg_id uuid.UUID) (order *order.Order, err error) {
	return s.repository.ReadOrderByMsgId(ctx, msg_id)
}

func (s *service) CreateOrder(ctx context.Context, order *order.Order) (err error) {
	return s.repository.CreateOrder(ctx, order)
}

func (s *service) UpdateOrder(ctx context.Context, id uuid.UUID, upFn func(oldOrder *order.Order) (*order.Order, error)) (*order.Order, error) {
	return s.repository.UpdateOrder(ctx, id, upFn)
}

func (s *service) DeleteOrderById(ctx context.Context, id uuid.UUID) (err error) {
	return s.repository.DeleteOrderById(ctx, id)
}
