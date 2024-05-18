package natsstreaming

import (
	"context"
	"encoding/json"

	"github.com/gogapopp/L0/internal/models"
	"github.com/nats-io/stan.go"
	"go.uber.org/zap"
)

type storager interface {
	AddOrder(context.Context, models.Order) error
}

func (n *nats) Sub(logger *zap.SugaredLogger, store storager) (stan.Subscription, error) {
	sub, err := n.conn.Subscribe("orders", func(m *stan.Msg) {
		var order models.Order

		err := json.Unmarshal(m.Data, &order)
		if err != nil {
			logger.Errorf("error unmarshalling message from nats: %w", err)
			return
		}

		logger.Infof("received an order: %+v", order)

		err = store.AddOrder(context.Background(), order)
		if err != nil {
			logger.Errorf("error saving order to database: %w", err)
			return
		}
	}, stan.DurableName("durable"))

	return sub, err
}
