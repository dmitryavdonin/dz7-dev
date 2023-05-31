package repository

import (
	"context"
	"errors"
	"order/internal/domain/message"
	"order/internal/repository/postgres/message/dao"
	"order/pkg/tools/transaction"
	"time"

	"github.com/jackc/pgx/v5"
)

const (
	tableName = "public.message"
)

var (
	ErrDuplicateKey = errors.New("ERROR: duplicate key value violates unique constraint \"message_pkey\" (SQLSTATE 23505)")
	ErrNotFound     = errors.New("not found")
	ErrUpdate       = errors.New("error updating or no changes")
	ErrEmptyResult  = errors.New("no rows in result set")
)

func (r *Repository) CreateMessage(ctx context.Context, message *message.Message) (err error) {
	tx, err := r.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	rawQuery := r.Builder.Insert(tableName).Columns(dao.OrderColumns...).Values(
		message.Id(), message.CreatedAt())
	query, args, err := rawQuery.ToSql()
	if err != nil {
		return
	}

	_, err = tx.Exec(ctx, query, args...)
	if err != nil {
		return
	}
	return
}

func (r *Repository) DeleteOldMessages(ctx context.Context, timeStamp time.Time) (err error) {
	tx, err := r.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	rawQuery := r.Builder.Delete(tableName).Where("createdAt <", timeStamp)
	query, args, err := rawQuery.ToSql()
	if err != nil {
		return
	}

	res, err := tx.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	if res.RowsAffected() == 0 {
		return ErrNotFound
	}
	return
}
