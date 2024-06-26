package cacherecoverer

import (
	"context"

	"github.com/gogapopp/L0/internal/models"
	"go.uber.org/zap"
)

type storager interface {
	GetAllOrders(ctx context.Context) ([]models.Order, error)
}

type cacher interface {
	SetOrderInCache(order models.Order)
}

func CacheRecover(logger *zap.SugaredLogger, cache cacher, store storager) {
	orders, err := store.GetAllOrders(context.Background())
	if err != nil {
		logger.Errorf("failed to get all orders from the database: %w", err)
	}
	for _, order := range orders {
		cache.SetOrderInCache(order)
	}
}
