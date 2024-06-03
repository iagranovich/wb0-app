package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"
	"wb0-app/config"
	"wb0-app/logger"
	m "wb0-app/models"
	"wb0-app/storage"

	"github.com/nats-io/stan.go"
)

func main() {
	fmt.Println("Hello world!")
	config.Load()
	logg := logger.Setup()
	// test connection to db
	s := storage.New()
	// test save order to db
	s.SaveOrder(&m.Order{OrderUid: "mnd0cm5sn7b2cms", TrackNumber: "WB0000000"},
		&m.Payment{OrderUid: "mnd0cm5sn7b2cms"},
		&m.Item{ChrtId: 99999, TrackNumber: "WB0000000"},
		&m.Delivery{OrderUid: "mnd0cm5sn7b2cms"})

	sub, err := stan.Connect(os.Getenv("BROKER_CID"), "sub")
	if err != nil {
		logg.Error("subscriber: cannot connect to NATS", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer sub.Close()

	_, err = sub.Subscribe("order-channel", func(m *stan.Msg) {
		fmt.Printf("Got: %s\n", string(m.Data))
	},
		stan.DurableName("dsub"))
	if err != nil {
		logg.Error("subscriber: cannot subscribe", slog.String("error", err.Error()))
	}

	for {
		time.Sleep(5 * time.Second)
	}

}
