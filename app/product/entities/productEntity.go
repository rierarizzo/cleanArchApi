package entities

import "time"

type (
	Product struct {
		Id           int             `json:"id,omitempty"`
		Category     ProductCategory `json:"category"`
		Name         string          `json:"name,omitempty"`
		Description  string          `json:"description,omitempty"`
		Price        float64         `json:"price,omitempty"`
		Cost         float64         `json:"cost,omitempty"`
		Quantity     int             `json:"quantity,omitempty"`
		Size         ProductSize     `json:"size"`
		Color        ProductColor    `json:"color"`
		Brand        string          `json:"brand,omitempty"`
		Sku          string          `json:"sku,omitempty"`
		Upc          string          `json:"upc,omitempty"`
		ImageUrl     string          `json:"image_url,omitempty"`
		Source       ProductSource   `json:"source"`
		SourceUrl    string          `json:"source_url,omitempty"`
		Offer        bool            `json:"offer,omitempty"`
		OfferPercent int             `json:"offer_percent,omitempty"`
		Active       bool            `json:"active,omitempty"`
		CreatedAt    time.Time       `json:"created_at"`
		UpdatedAt    time.Time       `json:"updated_at"`
	}

	ProductCategory struct {
		Id             int              `json:"id,omitempty"`
		ParentCategory *ProductCategory `json:"parent_category,omitempty"`
		Name           string           `json:"name,omitempty"`
		Description    string           `json:"description,omitempty"`
	}

	ProductSize struct {
		Code        string `json:"code,omitempty"`
		Description string `json:"description,omitempty"`
	}

	ProductColor struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Hex  string `json:"hex,omitempty"`
	}

	ProductSource struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	}

	ProductsOrder struct {
		Id            int       `json:"id,omitempty"`
		WeightPayment float64   `json:"weight_payment,omitempty"`
		Taxes         float64   `json:"taxes,omitempty"`
		Products      []Product `json:"products,omitempty"`
		TotalAmount   float64   `json:"total_amount,omitempty"`
		ArrivedAt     time.Time `json:"arrived_at"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
	}
)
