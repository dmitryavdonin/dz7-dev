package service

import (
	"order/internal/repository"
	order "order/internal/service/order"

	"github.com/dmitryavdonin/gtools/logger"
)

type Service struct {
	Order order.Service
}

func NewServices(repository *repository.Repository, logger logger.Interface) (*Service, error) {
	order, err := order.New(repository.Order, logger, order.Options{})
	if err != nil {
		return nil, err
	}

	return &Service{
		Order: order,
	}, nil
}
