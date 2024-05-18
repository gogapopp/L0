package natsstreaming

import (
	"github.com/gogapopp/L0/internal/config"
	"github.com/nats-io/stan.go"
)

type nats struct {
	conn stan.Conn
}

func Connect(config *config.Config) (*nats, error) {
	sc, err := stan.Connect(config.Stan.StanClusterID, config.Stan.ClientID, stan.NatsURL(config.Stan.DSN))
	if err != nil {
		return &nats{}, err
	}
	return &nats{conn: sc}, nil
}

func (n *nats) Close() {
	n.conn.Close()
}
