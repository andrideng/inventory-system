package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Product represents an products table.
type Product struct {
	SKU       string `json:"sku" db:"pk,sku"`
	Name      string `json:"name" db:"name"`
	Amount    int64  `json:"amount" db:"amount"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
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
	)
}
