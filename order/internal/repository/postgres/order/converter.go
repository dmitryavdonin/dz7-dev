package repository

import (
	"order/internal/domain/order"
	"order/internal/repository/postgres/order/dao"
)

func (r *Repository) toDomainOrder(daoOrder *dao.Order) (*order.Order, error) {
	return order.NewOrderWithId(daoOrder.Id, daoOrder.ProductId, daoOrder.ProductCount, daoOrder.ProductPrice, daoOrder.CreatedAt, daoOrder.ModifiedAt), nil
}
