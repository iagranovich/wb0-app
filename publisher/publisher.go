package main

import (
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/nats-io/stan.go"
)

func main() {
	pub, err := stan.Connect("wb0-app", "pub")
	if err != nil {
		slog.Error("publisher: cannot connect to NATS", slog.String("error", err.Error()))
		os.Exit(1)
	}
	slog.Info("publisher: connect to NATS successful")
	defer pub.Close()

	for i := 0; ; i++ {
		err = pub.Publish("order-channel", []byte("Order "+strconv.Itoa(i)))
		if err != nil {
			slog.Error("publisher: cannot publish", slog.String("error", err.Error()))
		}
		time.Sleep(1 * time.Second)
	}

}
