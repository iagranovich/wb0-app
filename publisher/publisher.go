package main

import (
	"fmt"
	"log/slog"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/nats-io/stan.go"
)

func randomString(length int, charset string) string {
	chars := []rune(charset)
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteRune(chars[rand.Intn(len(chars))])
	}
	return sb.String()
}

func randomOrder() string {
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	numbers := "0123456789"

	orderUid := randomString(19, lowercase+numbers)
	trackNumber := randomString(14, uppercase+numbers)
	chrtId, _ := strconv.Atoi(randomString(9, numbers))

	jsonOrder := fmt.Sprintf(`{
		"order_uid": %s,
		"track_number": %s,
		"entry": "WBIL",
		"delivery": {
		  "name": "Test Testov",
		  "phone": "+9720000000",
		  "zip": "2639809",
		  "city": "Kiryat Mozkin",
		  "address": "Ploshad Mira 15",
		  "region": "Kraiot",
		  "email": "test@gmail.com"
		},
		"payment": {
		  "transaction": "b563feb7b2b84b6test",
		  "request_id": "",
		  "currency": "USD",
		  "provider": "wbpay",
		  "amount": 1817,
		  "payment_dt": 1637907727,
		  "bank": "alpha",
		  "delivery_cost": 1500,
		  "goods_total": 317,
		  "custom_fee": 0
		},
		"items": [
		  {
			"chrt_id": %d,
			"track_number": %s,
			"price": 453,
			"rid": "ab4219087a764ae0btest",
			"name": "Mascaras",
			"sale": 30,
			"size": "0",
			"total_price": 317,
			"nm_id": 2389212,
			"brand": "Vivienne Sabo",
			"status": 202
		  }
		],
		"locale": "en",
		"internal_signature": "",
		"customer_id": "test",
		"delivery_service": "meest",
		"shardkey": "9",
		"sm_id": 99,
		"date_created": "2021-11-26T06:22:19Z",
		"oof_shard": "1"
	  }`, orderUid, trackNumber, chrtId, trackNumber)

	return jsonOrder
}

func main() {
	pub, err := stan.Connect("wb0-app", "pub")
	if err != nil {
		slog.Error("publisher: cannot connect to NATS", slog.String("error", err.Error()))
		os.Exit(1)
	}
	slog.Info("publisher: connect to NATS successful")
	defer pub.Close()

	order := randomOrder()
	for i := 0; ; i++ {
		err = pub.Publish("order-channel", []byte(order))
		if err != nil {
			slog.Error("publisher: cannot publish", slog.String("error", err.Error()))
		}
		time.Sleep(1 * time.Second)
	}

}
