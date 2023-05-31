package message

import (
	messageRepo "order/internal/repository/postgres/message"
	adaptersRepo "order/internal/service/message/adapters/repository"

	"github.com/dmitryavdonin/gtools/logger"
)

type service struct {
	repository adaptersRepo.Message
	logger     logger.Interface
	options    Options
}

type Options struct {
}

func (uc *service) SetOptions(options Options) {
	if uc.options != options {
		uc.options = options
	}
}

func New(repository *messageRepo.Repository, logger logger.Interface, options Options) (*service, error) {
	service := &service{
		repository: repository,
		logger:     logger,
	}

	service.SetOptions(options)
	return service, nil
}
