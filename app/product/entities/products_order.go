package entities

import "time"

type ProductsOrder struct {
	Id            int       `json:"id,omitempty"`
	WeightPayment float64   `json:"weight_payment,omitempty"`
	Taxes         float64   `json:"taxes,omitempty"`
	Products      []Product `json:"products,omitempty"`
	TotalAmount   float64   `json:"total_amount,omitempty"`
	ArrivedAt     time.Time `json:"arrived_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
