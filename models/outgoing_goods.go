package models

import validation "github.com/go-ozzo/ozzo-validation"

// OutgoingGoods represents an outgoing_goods table.
type OutgoingGoods struct {
	ID        int64   `json:"id" db:"pk,id"`
	SKU       string  `json:"sku" db:"sku"`
	Amount    int64   `json:"amount" db:"amount"`
	Price     float64 `json:"price" db:"price"`
	Total     float64 `json:"total" db:"total"`
	Note      string  `json:"note" db:"note"`
	CreatedAt string  `json:"created_at" db:"created_at"`
	UpdatedAt string  `json:"updated_at" db:"updated_at"`
}

// TableName match with database table name
func (m OutgoingGoods) TableName() string {
	return "outgoing_goods"
}

// Validate validate the OutgoingGoods Field
func (m OutgoingGoods) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.SKU, validation.Required, validation.Length(0, 100)),
		validation.Field(&m.Amount, validation.Required, validation.Min(0)),
		validation.Field(&m.Price, validation.Required, validation.Min(0.0)),
		validation.Field(&m.Note, validation.Required, validation.Length(0, 200)),
	)
}
