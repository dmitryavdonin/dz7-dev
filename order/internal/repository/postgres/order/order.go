package repository

import (
	"context"
	"errors"
	"order/internal/domain/order"
	"order/internal/repository/postgres/order/dao"
	"order/pkg/tools/transaction"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

const (
	tableName = "public.order"
)

var (
	ErrDuplicateKey = errors.New("ERROR: duplicate key value violates unique constraint \"order_pkey\" (SQLSTATE 23505)")
	ErrNotFound     = errors.New("not found")
	ErrUpdate       = errors.New("error updating or no changes")
	ErrEmptyResult  = errors.New("no rows in result set")
)

func (r *Repository) CreateOrder(ctx context.Context, order *order.Order) (err error) {
	tx, err := r.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	rawQuery := r.Builder.Insert(tableName).Columns(dao.OrderColumns...).Values(
		order.Id(),
		order.MsgId(),
		order.ProductId(),
		order.ProductCount(),
		order.ProductPrice(),
		order.Version(),
		order.CreatedAt(),
		order.ModifiedAt())
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

func (r *Repository) UpdateOrder(ctx context.Context, id uuid.UUID,
	upFunc func(oldOrder *order.Order) (*order.Order, error)) (order *order.Order, err error) {
	tx, err := r.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	oldOrder, err := r.oneOrderTx(ctx, tx, id, "id")
	if err != nil {
		return
	}

	newOrder, err := upFunc(oldOrder)
	if err != nil {
		return
	}

	rawQuery := r.Builder.Update(tableName).
		Set("product_id", newOrder.ProductId()).
		Set("product_count", newOrder.ProductCount()).
		Set("product_price", newOrder.ProductPrice()).
		Set("version", newOrder.Version()+1).
		Set("modified_at", newOrder.ModifiedAt()).Where("id = ?", newOrder.Id())
	query, args, err := rawQuery.ToSql()
	if err != nil {
		return
	}

	res, err := tx.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	if res.RowsAffected() == 0 {
		return nil, ErrUpdate
	}

	return newOrder, nil
}

func (r *Repository) DeleteOrderById(ctx context.Context, id uuid.UUID) (err error) {
	tx, err := r.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	rawQuery := r.Builder.Delete(tableName).Where("id = ?", id)
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

func (r *Repository) ReadOrderById(ctx context.Context, id uuid.UUID) (order *order.Order, err error) {
	tx, err := r.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	return r.oneOrderTx(ctx, tx, id, "id")
}

func (r *Repository) ReadOrderByMsgId(ctx context.Context, msg_id uuid.UUID) (order *order.Order, err error) {
	tx, err := r.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	return r.oneOrderTx(ctx, tx, msg_id, "msg_id")
}

func (r *Repository) oneOrderTx(ctx context.Context, tx pgx.Tx, field_value uuid.UUID, field_name string) (order *order.Order, err error) {
	rawQuery := r.Builder.Select(dao.OrderColumns...).From(tableName).Where(field_name+" = ?", field_value)
	query, args, err := rawQuery.ToSql()
	if err != nil {
		return
	}

	row, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	daoOrder, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[dao.Order])
	if err != nil {
		return nil, err
	}

	return r.toDomainOrder(&daoOrder)
}
