package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"
	"wb0-app/config"
	"wb0-app/logger"

	"github.com/nats-io/stan.go"
)

func main() {
	fmt.Println("App run!")
	config.Load()
	logg := logger.Setup()

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
