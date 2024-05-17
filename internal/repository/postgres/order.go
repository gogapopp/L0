package postgres

import "github.com/gogapopp/L0/internal/models"

func GetOrder(orderUID string) (models.Order, error) {
	const op = "postgres.order.GetOrder"
	return models.Order{}, nil
}

func AddOrder(models.Order) error {
	return nil
}
