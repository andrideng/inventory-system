package models

import validation "github.com/go-ozzo/ozzo-validation"

// IncomingGoods represnts an incoming_goods table.
type IncomingGoods struct {
	ID             int64   `json:"id" db:"pk,id"`
	ReceiptNumber  string  `json:"receipt_number" db:"receipt_number"`
	SKU            string  `json:"sku" db:"sku"`
	AmountOrders   int64   `json:"amount_orders" db:"amount_orders"`
	AmountReceived int64   `json:"amount_received" db:"amount_received"`
	PurchasePrice  float64 `json:"purchase_price" db:"purchase_price"`
	Total          float64 `json:"total" db:"total"`
	Note           string  `json:"note" db:"note"`
	Status         string  `json:"status" db:"status"`
	CreatedAt      string  `json:"created_at" db:"created_at"`
	UpdatedAt      string  `json:"updated_at" db:"updated_at"`
}

// TableName match with database table name
func (m IncomingGoods) TableName() string {
	return "incoming_goods"
}

// Validate validate the IncomingGoods fields
func (m IncomingGoods) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.ReceiptNumber, validation.Required, validation.Length(0, 100)),
		validation.Field(&m.SKU, validation.Required, validation.Length(0, 100)),
		validation.Field(&m.AmountOrders, validation.Required, validation.Min(0)),
		validation.Field(&m.AmountReceived, validation.Required, validation.Max(m.AmountOrders)),
		validation.Field(&m.PurchasePrice, validation.Required, validation.Min(0.0)),
	)
}

// ValidateUpdate validate update the IncomingGoods fields
func (m IncomingGoods) ValidateUpdate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.AmountReceived, validation.Required, validation.Min(0)),
	)
}
