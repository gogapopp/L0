package service

import (
	"context"
	"fmt"

	"github.com/gogapopp/L0/internal/models"
)

func (s *service) GetOrder(ctx context.Context, orderUID string) (models.Order, error) {
	const op = "service.order.GetOrder"

	order, ok := s.cache.GetOrderFromCache(orderUID)
	if ok {
		return order, nil
	}

	order, err := s.store.GetOrder(ctx, orderUID)
	if err != nil {
		return models.Order{}, fmt.Errorf("%s: %w", op, err)
	}

	s.cache.SetOrderInCache(order)

	return order, nil
}
func (s *service) AddOrder(ctx context.Context, order models.Order) error {
	const op = "service.order.AddOrder"

	err := s.store.AddOrder(ctx, order)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	s.cache.SetOrderInCache(order)

	return nil
}
func (s *service) GetAllOrders(ctx context.Context) ([]models.Order, error) {
	const op = "service.order.GetAllOrders"

	orders, err := s.store.GetAllOrders(ctx)
	if err != nil {
		return []models.Order{}, fmt.Errorf("%s: %w", op, err)
	}

	return orders, nil
}
