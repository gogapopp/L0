package cacherecoverer

import (
	"github.com/gogapopp/L0/internal/models"
	"go.uber.org/zap"
)

type storager interface {
	GetAllOrders() ([]models.Order, error)
}

type cacher interface {
	SetOrderInCache(order models.Order)
}

func CacheRecover(logger *zap.SugaredLogger, cache cacher, store storager) {
	orders, err := store.GetAllOrders()
	logger.Error("failed to get all orders from the database: %w", err)
	for _, order := range orders {
		cache.SetOrderInCache(order)
	}
}
