package postgres

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gogapopp/L0/internal/models"
)

func (r *repository) GetOrder(ctx context.Context, orderUID string) (models.Order, error) {
	const (
		op    = "postgres.order.GetOrder"
		query = "SELECT data FROM orders WHERE data->>'order_uid' = $1"
	)

	row := r.db.QueryRow(ctx, query, orderUID)

	var (
		order models.Order
		data  []byte
	)

	if err := row.Scan(&data); err != nil {
		return models.Order{}, fmt.Errorf("%s: %w", op, err)
	}

	if err := json.Unmarshal(data, &order); err != nil {
		return models.Order{}, fmt.Errorf("%s: %w", op, err)
	}

	return order, nil
}

func (r *repository) AddOrder(ctx context.Context, order models.Order) error {
	const (
		op    = "postgres.order.AddOrder"
		query = "INSERT INTO orders (data) VALUES ($1)"
	)

	data, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = r.db.Exec(ctx, query, data)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *repository) GetAllOrders(ctx context.Context) ([]models.Order, error) {
	const (
		op    = "postgres.order.GetAllOrders"
		query = "SELECT data FROM orders"
	)

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var (
			order models.Order
			data  []byte
		)

		if err := rows.Scan(&data); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		if err := json.Unmarshal(data, &order); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return orders, nil
}
