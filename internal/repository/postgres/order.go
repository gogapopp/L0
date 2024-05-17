package postgres

import (
	"context"

	"github.com/gogapopp/L0/internal/models"
)

func (s *repository) GetOrder(ctx context.Context, orderUID string) (models.Order, error) {
	const op = "postgres.order.GetOrder"
	return models.Order{}, nil
}

func (s *repository) AddOrder(ctx context.Context, order models.Order) error {
	return nil
}

func (s *repository) GetAllOrders(ctx context.Context) ([]models.Order, error) {
	return []models.Order{}, nil
}
