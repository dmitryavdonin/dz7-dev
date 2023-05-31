package delivery

import (
	"context"
	"fmt"
	"net/http"
	jsonRequests "order/internal/delivery/http/order"
	"order/internal/domain/message"
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
	}
	fmt.Println("CreateOrder(): SUCCESS! Idempotency key = " + idempotencyKey)

	msg_id, err := uuid.Parse(idempotencyKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("CreateOrder(): FAILED! Cannot parse UUID from idempotency key, err = " + err.Error())
		return
	}

	message, err := message.NewMessage(msg_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("CreateOrder(): FAILED! Message cannot be created, err = " + err.Error())
		return
	}

	err = d.services.Message.CreateMessage(context.Background(), message)
	if err != nil {
		exist, err := d.services.Order.ReadOrderByMsgId(context.Background(), msg_id)
		if exist != nil {
			c.JSON(http.StatusCreated, d.toResponseOrder(exist))
			fmt.Println("CreateOrder(): SUCCESS! Order already exist id = " + exist.Id().String())
			return
		}

		if err != nil {
			fmt.Println("CreateOrder(): FAILED! err = " + err.Error())
		}
	}

	fmt.Println("CreateOrder(): SUCCESS! Message is created id = " + message.Id().String())

	request := jsonRequests.CreateOrderRequest{}

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("CreateOrder(): FAILED! Cannot bind parameters, err = " + err.Error())
		return
	}

	order, err := order.NewOrder(msg_id, request.ProductId, request.ProductCount, request.ProductPrice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = d.services.Order.CreateOrder(context.Background(), order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("CreateOrder(): FAILED! Cannot create order, err = " + err.Error())
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
		return order.NewOrderWithId(oldOrder.Id(), oldOrder.MsgId(), request.ProductId, request.ProductCount, request.ProductPrice, oldOrder.CreatedAt(), time.Now()), nil
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
