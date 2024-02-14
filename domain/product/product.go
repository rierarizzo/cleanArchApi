package product

import "time"

type (
	Product struct {
		Id           int         `json:"id,omitempty"`
		Category     Category    `json:"category"`
		Subcategory  Subcategory `json:"subcategory"`
		Name         string      `json:"name,omitempty"`
		Description  *string     `json:"description,omitempty"`
		Price        float64     `json:"price,omitempty"`
		Cost         float64     `json:"cost,omitempty"`
		Quantity     int         `json:"quantity,omitempty"`
		Size         Size        `json:"size"`
		Color        Color       `json:"color"`
		Brand        string      `json:"brand,omitempty"`
		Sku          string      `json:"sku,omitempty"`
		Upc          string      `json:"upc,omitempty"`
		ImageUrl     string      `json:"image_url,omitempty"`
		Source       Source      `json:"source"`
		SourceUrl    *string     `json:"source_url,omitempty"`
		IsOffered    bool        `json:"is_offered,omitempty"`
		OfferPercent *int        `json:"offer_percent,omitempty"`
		IsActive     bool        `json:"is_active,omitempty"`
		CreatedAt    time.Time   `json:"created_at"`
		UpdatedAt    time.Time   `json:"updated_at"`
	}

	Category struct {
		Id          int    `json:"id,omitempty"`
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
	}

	Subcategory struct {
		Id             int       `json:"id,omitempty"`
		ParentCategory *Category `json:"parent_category,omitempty"`
		Name           string    `json:"name,omitempty"`
		Description    string    `json:"description,omitempty"`
	}

	Color struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Hex  string `json:"hex,omitempty"`
	}

	Size struct {
		Code        string `json:"code,omitempty"`
		Description string `json:"description,omitempty"`
	}

	Source struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	}
)
