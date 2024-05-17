package natsstreaming

import (
	"github.com/nats-io/stan.go"
)

type nats struct {
	conn stan.Conn
}

func Connect() (*nats, error) {
	sc, err := stan.Connect("test-cluster", "client-id", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		return &nats{}, err
	}
	return &nats{conn: sc}, nil
}

func (n *nats) Close() {
	n.conn.Close()
	n.conn.NatsConn().Drain()
}
