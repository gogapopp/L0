package handler

import (
	"net/http"

	"go.uber.org/zap"
)

type servicer interface {
}

func GetOrderById(log *zap.SugaredLogger, service servicer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		_ = id
	}
}
