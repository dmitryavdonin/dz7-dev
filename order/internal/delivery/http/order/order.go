package order

import (
	"time"

	"github.com/google/uuid"
)

type CreateOrderRequest struct {
	ProductId    int     `json:"product_id" default:"0"`
	ProductCount int     `json:"product_count" default:"0"`
	ProductPrice float32 `json:"product_price" default:"0"`
}

type UpdateOrderRequest struct {
	ProductId    int     `json:"product_id" default:"1"`
	ProductCount int     `json:"product_count" default:"1"`
	ProductPrice float32 `json:"product_price" default:"1"`
}

type UpdateOrderResponse struct {
	Result OrderResponse `json:"result"`
}

type OrderResponse struct {
	Id           uuid.UUID `json:"id"`
	ProductId    int       `json:"product_id"`
	ProductCount int       `json:"product_count"`
	ProductPrice float32   `json:"product_price"`
	CreatedAt    time.Time `json:"created_at"`
	ModifiedAt   time.Time `json:"modified_at"`
}
