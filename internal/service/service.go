package service

import (
	"context"

	"github.com/gogapopp/L0/internal/models"
)

type service struct {
	store storager
	cache cacher
}

type cacher interface {
	GetOrderFromCache(orderUID string) (models.Order, bool)
	SetOrderInCache(order models.Order)
}

type storager interface {
	GetOrder(ctx context.Context, orderUID string) (models.Order, error)
	AddOrder(ctx context.Context, order models.Order) error
	GetAllOrders(ctx context.Context) ([]models.Order, error)
}

func New(store storager, cache cacher) *service {
	return &service{
		store: store,
		cache: cache,
	}
}
