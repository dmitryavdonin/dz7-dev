package delivery

import (
	"context"
	"fmt"
	"net/http"
	jsonRequests "order/internal/delivery/http/order"
	"order/internal/domain/order"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (d *Delivery) CreateOrder(c *gin.Context) {
	idempotencyKey := c.GetHeader("x-idempotency-key")
	if idempotencyKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Idempotency key is not presented"})
		fmt.Println("CreateOrder(): FAILED! Idempotency key is not presented")
		return
	} else {
		fmt.Println("CreateOrder(): Idempotency key= " + idempotencyKey)
	}

	request := jsonRequests.CreateOrderRequest{}

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := order.NewOrder(request.ProductId, request.ProductCount, request.ProductPrice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = d.services.Order.CreateOrder(context.Background(), order)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, d.toResponseOrder(order))
}

func (d *Delivery) UpdateOrder(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request := jsonRequests.UpdateOrderRequest{}

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	upFn := func(oldOrder *order.Order) (*order.Order, error) {
		return order.NewOrderWithId(oldOrder.Id(), request.ProductId, request.ProductCount, request.ProductPrice, oldOrder.CreatedAt(), time.Now()), nil
	}

	order, err := d.services.Order.UpdateOrder(context.Background(), id, upFn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, d.toResponseOrder(order))
}

func (d *Delivery) DeleteOrderById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = d.services.Order.DeleteOrderById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (d *Delivery) ReadOrderById(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := d.services.Order.ReadOrderById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, d.toResponseOrder(order))
}

func (d *Delivery) GetIdempotencyKey(c *gin.Context) {
	//c.JSON(http.StatusOK, uuid.NewString())
	c.String(http.StatusOK, uuid.NewString())
}
