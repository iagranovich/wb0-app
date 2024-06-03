package models

type Payment struct {
	OrderUid     string `db:"order_uid"`
	Transaction  string `db:"transaction"`
	RequestId    string `db:"request_id"`
	Currency     string `db:"currency"`
	Provider     string `db:"provider"`
	Amount       int    `db:"amount"`
	PaymentDt    int    `db:"payment_dt"`
	Bank         string `db:"bank"`
	DeliveryCost int    `db:"delivery_cost"`
	GoodsTotal   int    `db:"goods_total"`
	CustomFee    int    `db:"custom_fee"`
}
