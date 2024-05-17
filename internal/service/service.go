package service

type service struct {
	store storager
	cache cacher
}

type cacher interface {
}

type storager interface {
}

func New(store storager, cache cacher) *service {
	return &service{
		store: store,
		cache: cache,
	}
}
