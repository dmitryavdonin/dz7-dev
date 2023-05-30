package user

import (
	orderRepo "order/internal/repository/postgres/order"
	adaptersRepo "order/internal/service/order/adapters/repository"
	"time"

	"github.com/dmitryavdonin/gtools/logger"
)

type service struct {
	repository adaptersRepo.Order
	logger     logger.Interface
	//tokenManager tokenManager.TokenManager
	options Options
}

type Options struct {
	TokenTTL time.Duration
}

func (uc *service) SetOptions(options Options) {
	if uc.options != options {
		uc.options = options
	}
}

func New(repository *orderRepo.Repository, logger logger.Interface, options Options) (*service, error) {
	service := &service{
		repository: repository,
		logger:     logger,
	}

	service.SetOptions(options)
	return service, nil
}
