package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gogapopp/L0/internal/models"
	"github.com/gogapopp/L0/internal/repository"
	"go.uber.org/zap"
)

type servicer interface {
	GetOrder(ctx context.Context, orderUID string) (models.Order, error)
}

func GetOrderById(logger *zap.SugaredLogger, service servicer) http.HandlerFunc {
	const op = "handler.GetOrderById"
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		order, err := service.GetOrder(r.Context(), id)
		if err != nil {
			logger.Errorf("%s: %w", op, err)
			if errors.Is(err, repository.ErrOrderNotExist) {
				http.Error(w, "bad request", http.StatusBadRequest)
				return
			}
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(order); err != nil {
			logger.Errorf("%s: %w", op, err)
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}
	}
}
