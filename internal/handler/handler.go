package handler

import (
	"context"
	"errors"
	"html/template"
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
				http.Error(w, "order does not exist", http.StatusBadRequest)
				return
			}
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("templates/order_template.html")
		if err != nil {
			logger.Errorf("%s: %w", op, err)
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		type PageData struct {
			Order models.Order
		}

		data := PageData{
			Order: order,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			logger.Errorf("%s: %w", op, err)
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}
	}
}
