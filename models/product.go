package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Product represents an products table.
type Product struct {
	SKU                  string  `json:"sku" db:"pk,sku"`
	Name                 string  `json:"name" db:"name"`
	Size                 string  `json:"size" db:"size"`
	Color                string  `json:"color" db:"color"`
	Amount               int64   `json:"amount" db:"amount"`
	TotalIncomingAmount  int64   `json:"total_incoming_amount" db:"total_incoming_amount"`
	TotalIncomingPrice   float64 `json:"total_incoming_price" db:"total_incoming_price"`
	AveragePurchasePrice int64   `json:"average_purchase_price" db:"average_purchase_price"`
	CreatedAt            string  `json:"created_at" db:"created_at"`
	UpdatedAt            string  `json:"updated_at" db:"updated_at"`
}

// TableName match with database table name
func (m Product) TableName() string {
	return "products"
}

// Validate validates the Product fields
func (m Product) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.SKU, validation.Required, validation.Length(0, 100)),
		validation.Field(&m.Name, validation.Required, validation.Length(0, 100)),
		validation.Field(&m.Size, validation.Required, validation.Length(0, 100)),
		validation.Field(&m.Color, validation.Required, validation.Length(0, 100)),
	)
}
