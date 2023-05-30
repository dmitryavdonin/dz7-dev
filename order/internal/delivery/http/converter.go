package delivery

import (
	response "order/internal/delivery/http/order"
	"order/internal/domain/order"
)

func (d *Delivery) toResponseOrder(order *order.Order) *response.OrderResponse {
	return &response.OrderResponse{
		Id:           order.Id(),
		ProductId:    order.ProductId(),
		ProductCount: order.ProductCount(),
		ProductPrice: order.ProductPrice(),
		CreatedAt:    order.CreatedAt(),
		ModifiedAt:   order.ModifiedAt(),
	}
}
