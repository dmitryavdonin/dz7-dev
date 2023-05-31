package repository

import (
	"github.com/dmitryavdonin/gtools/psql"
)

type Repository struct {
	*psql.Postgres
}

type Options struct {
	DefaultLimit uint64
}

func New(pg *psql.Postgres, options Options) (*Repository, error) {
	var r = &Repository{pg}
	r.SetOptions(options)
	return r, nil
}

func (r *Repository) SetOptions(options Options) {
	if options.DefaultLimit == 0 {
		options.DefaultLimit = 10
	}
}
