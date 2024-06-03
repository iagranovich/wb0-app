package models

type Order struct {
	OrderUid          string `db:"order_uid"`
	TrackNumber       string `db:"track_number"`
	Entry             string `db:"entry"`
	Locale            string `db:"locale"`
	InternalSignature string `db:"internal_signature"`
	CustomerId        string `db:"customer_id"`
	DeliveryService   string `db:"delivery_service"`
	Shardkey          int    `db:"shardkey"`
	SmId              int    `db:"sm_id"`
	DateCreated       string `db:"date_created"`
	OofShard          int    `db:"oof_shard"`
}
