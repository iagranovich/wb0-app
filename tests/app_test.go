package hand_test

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
	"wb0-app/cache"
	"wb0-app/client"
	"wb0-app/models"
	"wb0-app/storage"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestSaveOrderToDB(t *testing.T) {

	storage := Init()

	// json to Order struct
	orderJson := `{
		"order_uid": "b563feb7b2b84b6test",
		"track_number": "WBILMTESTTRACK",
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
			"chrt_id": 9934930,
			"track_number": "WBILMTESTTRACK",
			"price": 453,
			"rid": "ab4219087a764ae0btest",
			"name": "Mascaras",
			"sale": 30,
			"size": "0",
			"total_price": 317,
			"nm_id": 2389212,
			"brand": "Vivienne Sabo",
			"status": 202
		  },
		  {
			"chrt_id": 9934931,
			"track_number": "WBILMTESTTRACK",
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
	  }`

	order := models.Order{}
	if err := json.Unmarshal([]byte(orderJson), &order); err != nil {
		log.Fatalf(err.Error())
	}

	storage.Save(order)
}

func TestStanClient_Subscribe(t *testing.T) {
	storage := Init()
	cache := cache.New()

	subscriber := client.New()
	subscriber.Subscribe(models.Order{}, storage.Save, cache.Save)

	time.Sleep(5 * time.Second)
}

func TestCache_Save_FindByUid(t *testing.T) {
	cache := cache.New()
	orderuid := "orderuid"
	order := models.Order{OrderUid: orderuid}

	cache.Save(order)

	_, err := cache.FindByUid("ssssssssss")
	require.Error(t, err)

	result, err := cache.FindByUid(orderuid)
	require.NoError(t, err)
	require.Equal(t, order, result)

}

type mockStorage struct{}

func (s *mockStorage) FindAll() []models.Order {
	orders := []models.Order{
		{
			OrderUid: "1111",
		},
		{
			OrderUid: "2222",
		},
	}

	return orders
}

func TestCache_Restore(t *testing.T) {
	ms := &mockStorage{}
	cache := cache.New()

	cache.Restore(ms)

	result, err := cache.FindByUid("1111")
	require.NoError(t, err)
	require.Equal(t, "1111", result.OrderUid)

	result, err = cache.FindByUid("2222")
	require.NoError(t, err)
	require.Equal(t, "2222", result.OrderUid)

	result, err = cache.FindByUid("3333")
	require.Error(t, err)
	require.Empty(t, result)

}

func Init() *storage.Storage {
	// config
	if err := godotenv.Load("test.env"); err != nil {
		log.Fatal(err.Error())
	}

	// connection to test db
	s := storage.New()

	// prepare test db
	m, err := migrate.New(
		"file://../migrations",
		fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable",
			os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE")))
	if err != nil {
		log.Fatal("op: migrate.new ", err.Error())
	}
	if err := m.Down(); err != nil {
		log.Fatal("op: m.down ", err.Error())
	}
	if err := m.Up(); err != nil {
		log.Fatal("op: m.up ", err.Error())
	}

	return s
}
