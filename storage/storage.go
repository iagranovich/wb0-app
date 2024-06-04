package storage

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	m "wb0-app/models"
)

type Storage struct {
	db *sqlx.DB
}

func New() (*Storage, error) {
	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_DATABASE"),
			os.Getenv("DB_PORT")))
	if err != nil {
		slog.Error("storage: cannot connect to db", slog.String("error", err.Error()))
		return nil, err
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveOrder(o *m.Order, p *m.Payment, i []m.Item, d *m.Delivery) {
	qO := `INSERT INTO orders (
		    order_uid, track_number, entry, locale,
		    internal_signature, customer_id, delivery_service, shardkey,
		    sm_id, date_created, oof_shard
		)
		VALUES (
			:order_uid, :track_number, :entry, :locale,
			:internal_signature, :customer_id, :delivery_service,
			:shardkey, :sm_id, :date_created, :oof_shard
		)`
	qP := `INSERT INTO payments (
		    order_uid, transaction, request_id, currency, provider,
			amount, payment_dt, bank, delivery_cost, goods_total,
			custom_fee
	    )
		VALUES (
            :order_uid, :transaction, :request_id, :currency, :provider,
			:amount, :payment_dt, :bank, :delivery_cost, :goods_total,
			:custom_fee
		)`
	qI := `INSERT INTO items (
		    chrt_id, track_number, price, rid, name, sale, 
			size, total_price, nm_id, brand, status
	    )
	    VALUES (
			:chrt_id, :track_number, :price, :rid, :name, :sale, 
			:size, :total_price, :nm_id, :brand, :status
		)`
	qD := `INSERT INTO deliveries (
		    order_uid, name, phone, zip, city, address, region, email
	    )
		VALUES (
			:order_uid, :name, :phone, :zip, :city, :address, :region, :email
		)`

	tx := s.db.MustBegin()
	if _, err := tx.NamedExec(qO, o); err != nil {
		slog.Error("storage: cannot save data to Orders table", slog.String("error", err.Error()))
	}
	if _, err := tx.NamedExec(qP, p); err != nil {
		slog.Error("storage: cannot save data to Payments table", slog.String("error", err.Error()))
	}
	if _, err := tx.NamedExec(qI, i); err != nil {
		slog.Error("storage: cannot save data to Items table", slog.String("error", err.Error()))
	}
	if _, err := tx.NamedExec(qD, d); err != nil {
		slog.Error("storage: cannot save data to Deliveries table", slog.String("error", err.Error()))
	}
	tx.Commit()
}
