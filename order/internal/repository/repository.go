package repository

import (
	order "order/internal/repository/postgres/order"

	"github.com/dmitryavdonin/gtools/psql"
)

type Repository struct {
	Order *order.Repository
}

func NewRepository(pg *psql.Postgres) (*Repository, error) {
	order, err := order.New(pg, order.Options{})
	if err != nil {
		return nil, err
	}

	return &Repository{
		Order: order,
	}, nil
}
