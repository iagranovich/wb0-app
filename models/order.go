package models

type Order struct {
	OrderUid          string   `db:"order_uid" json:"order_uid"`
	TrackNumber       string   `db:"track_number" json:"track_number"`
	Entry             string   `db:"entry" json:"entry"`
	Locale            string   `db:"locale" json:"locale"`
	InternalSignature string   `db:"internal_signature" json:"internal_signature"`
	CustomerId        string   `db:"customer_id" json:"customer_id"`
	DeliveryService   string   `db:"delivery_service" json:"delivery_service"`
	Shardkey          string   `db:"shardkey" json:"shardkey"`
	SmId              int      `db:"sm_id" json:"sm_id"`
	DateCreated       string   `db:"date_created" json:"date_created"`
	OofShard          string   `db:"oof_shard" json:"oof_shard"`
	Payment           Payment  `json:"payment"`
	Items             []Item   `json:"items"`
	Delivery          Delivery `json:"delivery"`
}
