package repository

import (
	message "order/internal/repository/postgres/message"
	order "order/internal/repository/postgres/order"

	"github.com/dmitryavdonin/gtools/psql"
)

type Repository struct {
	Order   *order.Repository
	Message *message.Repository
}

func NewRepository(pg *psql.Postgres) (*Repository, error) {
	order, err := order.New(pg, order.Options{})
	if err != nil {
		return nil, err
	}

	message, err := message.New(pg, message.Options{})
	if err != nil {
		return nil, err
	}

	return &Repository{
		Order:   order,
		Message: message,
	}, nil
}
