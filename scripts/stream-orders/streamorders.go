package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/gogapopp/L0/internal/config"
	"github.com/gogapopp/L0/internal/libs/logger"
	"github.com/gogapopp/L0/internal/models"
	"github.com/nats-io/stan.go"
)

type nats struct {
	conn stan.Conn
}

func main() {
	logger := must(logger.New())
	config := must(config.New())
	conn := must(connect(config))
	defer conn.close()

	ordersNums := 1

	orders := createOrders(ordersNums)

	for _, order := range orders {
		err := conn.publishOrder(order)
		if err != nil {
			logger.Errorf("error publish order: %w", err)
		}
	}
}

func (n *nats) publishOrder(order models.Order) error {
	data, err := json.Marshal(order)
	if err != nil {
		return err
	}
	return n.conn.Publish("orders", data)
}

func connect(config *config.Config) (*nats, error) {
	sc, err := stan.Connect(config.Stan.StanClusterID, fmt.Sprintf("%s%d", config.Stan.ClientID, 1), stan.NatsURL(config.Stan.DSN))
	if err != nil {
		return &nats{}, err
	}
	return &nats{conn: sc}, nil
}

func (n *nats) close() {
	n.conn.Close()
}

func createOrders(ordersNums int) []models.Order {
	var orders []models.Order
	gofakeit.Seed(time.Now().UnixNano())

	for i := 0; i < ordersNums; i++ {
		order := models.Order{
			OrderUID:    gofakeit.UUID(),
			TrackNumber: gofakeit.UUID(),
			Entry:       gofakeit.Sentence(5),
			Delivery: models.Delivery{
				Name:    gofakeit.Name(),
				Phone:   gofakeit.Phone(),
				Zip:     gofakeit.Zip(),
				City:    gofakeit.City(),
				Address: gofakeit.Address().Address,
				Region:  gofakeit.State(),
				Email:   gofakeit.Email(),
			},
			Payment: models.Payment{
				Transaction:  gofakeit.UUID(),
				RequestID:    gofakeit.UUID(),
				Currency:     "USD",
				Provider:     "Bank",
				Amount:       gofakeit.Number(100, 10000),
				PaymentDt:    gofakeit.Date().Year(),
				Bank:         gofakeit.Company(),
				DeliveryCost: gofakeit.Number(100, 500),
				GoodsTotal:   gofakeit.Number(500, 5000),
				CustomFee:    gofakeit.Number(50, 200),
			},
			Items: []models.Item{
				{
					ChrtID:      gofakeit.Number(1000, 9999),
					TrackNumber: gofakeit.UUID(),
					Price:       gofakeit.Number(100, 1000),
					RID:         gofakeit.UUID(),
					Name:        gofakeit.Name(),
					Sale:        gofakeit.Number(10, 50),
					Size:        gofakeit.UUID(),
					TotalPrice:  gofakeit.Number(100, 1000),
					NmID:        gofakeit.Number(1000, 9999),
					Brand:       gofakeit.Vehicle().Brand,
					Status:      gofakeit.Number(1, 5),
				},
			},
			Locale:            gofakeit.City(),
			InternalSignature: gofakeit.UUID(),
			CustomerID:        gofakeit.UUID(),
			DeliveryService:   gofakeit.Company(),
			ShardKey:          gofakeit.UUID(),
			SmID:              gofakeit.Number(1000, 9999),
			DateCreated:       gofakeit.Date(),
			OofShard:          gofakeit.UUID(),
		}

		orders = append(orders, order)
		fmt.Printf("Order %s\n", order.OrderUID)
	}
	return orders
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
