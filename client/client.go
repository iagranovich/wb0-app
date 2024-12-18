package client

import (
	"encoding/json"
	"log/slog"
	"os"
	"wb0-app/models"

	"github.com/nats-io/stan.go"
)

type stanClient struct {
	Conn stan.Conn
}

func New() Client {
	conn, err := stan.Connect(os.Getenv("BROKER_CID"), "sub")
	if err != nil {
		slog.Error("subscriber: cannot conect to NATS", slog.String("error", err.Error()))
		os.Exit(1)
	}
	return &StanClient{Conn: conn}
}

func (sc stanClient) Subscribe(order models.Order, orderHandler ...func(models.Order)) {

	_, err := sc.Conn.Subscribe("order-channel", func(m *stan.Msg) {

		err := json.Unmarshal(m.Data, &order)
		if err != nil {
			slog.Error("subscriber: cannot unmarshal order", slog.String("error", err.Error()))
			return
		}

		for _, handler := range orderHandler {
			handler(order)
		}
	},
		stan.DurableName("dsub"))
	if err != nil {
		slog.Error("subscriber: cannot subscribe", slog.String("error", err.Error()))
		os.Exit(1)
	}

}
