package entities

import "time"

type (
	Product struct {
		Id           int
		Category     ProductCategory
		Name         string
		Description  string
		Price        float64
		Cost         float64
		quantity     int
		Size         ProductSize
		Color        ProductColor
		Brand        string
		Sku          string
		Upc          string
		ImageUrl     string
		Source       ProductSource
		SourceUrl    string
		Offer        bool
		OfferPercent int
		Active       bool
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}

	ProductCategory struct {
		Id             int
		ParentCategory *ProductCategory
		Name           string
		Description    string
	}

	ProductSize struct {
		Code        string
		Description string
	}

	ProductColor struct {
		Id   int
		Name string
		Hex  string
	}

	ProductSource struct {
		Id   int
		Name string
	}

	ProductsOrder struct {
		Id            int
		WeightPayment float64
		Taxes         float64
		Products      []Product
		TotalAmount   float64
		ArrivedAt     time.Time
		CreatedAt     time.Time
		UpdatedAt     time.Time
	}
)
