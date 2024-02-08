package product

import "time"

type Product struct {
	Id           int                `json:"id,omitempty"`
	Category     ProductCategory    `json:"category"`
	Subcategory  ProductSubcategory `json:"subcategory"`
	Name         string             `json:"name,omitempty"`
	Description  *string            `json:"description,omitempty"`
	Price        float64            `json:"price,omitempty"`
	Cost         float64            `json:"cost,omitempty"`
	Quantity     int                `json:"quantity,omitempty"`
	Size         ProductSize        `json:"size"`
	Color        ProductColor       `json:"color"`
	Brand        string             `json:"brand,omitempty"`
	Sku          string             `json:"sku,omitempty"`
	Upc          string             `json:"upc,omitempty"`
	ImageUrl     string             `json:"image_url,omitempty"`
	Source       ProductSource      `json:"source"`
	SourceUrl    *string            `json:"source_url,omitempty"`
	IsOffered    bool               `json:"is_offered,omitempty"`
	OfferPercent *int               `json:"offer_percent,omitempty"`
	IsActive     bool               `json:"is_active,omitempty"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
}
