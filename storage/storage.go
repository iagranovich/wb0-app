package storage

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	m "wb0-app/models"
)

type storage struct {
	db *sqlx.DB
}

func New() Storage {
	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_DATABASE"),
			os.Getenv("DB_PORT")))
	if err != nil {
		slog.Error("storage: cannot connect to db", slog.String("error", err.Error()))
		os.Exit(1)
	}

	return &storage{db: db}
}

func (s storage) Save(order m.Order) {

	payment := &order.Payment
	delivery := &order.Delivery
	items := order.Items

	payment.OrderUid = order.OrderUid
	delivery.OrderUid = order.OrderUid

	orderQuery := `INSERT INTO orders (
		    order_uid, track_number, entry, locale,
		    internal_signature, customer_id, delivery_service, shardkey,
		    sm_id, date_created, oof_shard
		)
		VALUES (
			:order_uid, :track_number, :entry, :locale,
			:internal_signature, :customer_id, :delivery_service,
			:shardkey, :sm_id, :date_created, :oof_shard
		)`
	paymentQuery := `INSERT INTO payments (
		    order_uid, transaction, request_id, currency, provider,
			amount, payment_dt, bank, delivery_cost, goods_total,
			custom_fee
	    )
		VALUES (
            :order_uid, :transaction, :request_id, :currency, :provider,
			:amount, :payment_dt, :bank, :delivery_cost, :goods_total,
			:custom_fee
		)`
	itemQuery := `INSERT INTO items (
		    chrt_id, track_number, price, rid, name, sale, 
			size, total_price, nm_id, brand, status
	    )
	    VALUES (
			:chrt_id, :track_number, :price, :rid, :name, :sale, 
			:size, :total_price, :nm_id, :brand, :status
		)`
	deliveryQuery := `INSERT INTO deliveries (
		    order_uid, name, phone, zip, city, address, region, email
	    )
		VALUES (
			:order_uid, :name, :phone, :zip, :city, :address, :region, :email
		)`

	tx := s.db.MustBegin()
	if _, err := tx.NamedExec(orderQuery, order); err != nil {
		slog.Error("storage: cannot save data to Orders table", slog.String("error", err.Error()))
	}
	if _, err := tx.NamedExec(paymentQuery, payment); err != nil {
		slog.Error("storage: cannot save data to Payments table", slog.String("error", err.Error()))
	}
	if _, err := tx.NamedExec(itemQuery, items); err != nil {
		slog.Error("storage: cannot save data to Items table", slog.String("error", err.Error()))
	}
	if _, err := tx.NamedExec(deliveryQuery, delivery); err != nil {
		slog.Error("storage: cannot save data to Deliveries table", slog.String("error", err.Error()))
	}
	tx.Commit()
}

func (s storage) FindAll() []m.Order {
	orders := []m.Order{}
	s.db.Select(&orders, "SELECT * FROM orders")

	var result []m.Order
	for _, order := range orders {
		items := []m.Item{}
		s.db.Select(&items, "SELECT * FROM items WHERE track_number=$1", order.TrackNumber)
		order.Items = items

		payment := m.Payment{}
		s.db.Get(&payment, "SELECT * FROM payments WHERE order_uid=$1", order.OrderUid)
		order.Payment = payment

		delivery := m.Delivery{}
		s.db.Get(&delivery, "SELECT * FROM deliveries WHERE order_uid=$1", order.OrderUid)
		order.Delivery = delivery

		result = append(result, order)
	}

	return result
}
