package service

import (
	"order/internal/repository"
	message "order/internal/service/message"
	order "order/internal/service/order"

	"github.com/dmitryavdonin/gtools/logger"
)

type Service struct {
	Order   order.Service
	Message message.Service
}

func NewServices(repository *repository.Repository, logger logger.Interface) (*Service, error) {

	order, err := order.New(repository.Order, logger, order.Options{})
	if err != nil {
		return nil, err
	}

	message, err := message.New(repository.Message, logger, message.Options{})
	if err != nil {
		return nil, err
	}

	return &Service{
		Order:   order,
		Message: message,
	}, nil
}
