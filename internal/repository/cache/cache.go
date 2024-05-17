package cache

import (
	"time"

	"github.com/gogapopp/L0/internal/models"
	"github.com/patrickmn/go-cache"
)

type cacheRepo struct {
	cache *cache.Cache
}

func New(defaultTTL, cleanupInterval time.Duration) *cacheRepo {
	cache := cache.New(defaultTTL, cleanupInterval)
	return &cacheRepo{cache: cache}
}

func (c *cacheRepo) GetOrderFromCache(orderUID string) (models.Order, bool) {
	orderCache, exist := c.cache.Get(orderUID)
	if !exist {
		return models.Order{}, false
	}

	order, ok := orderCache.(models.Order)
	if !ok {
		return models.Order{}, false
	}

	return order, true
}

func (c *cacheRepo) SetOrderInCache(order models.Order) {
	c.cache.Set(order.OrderUID, order, 0)
}
