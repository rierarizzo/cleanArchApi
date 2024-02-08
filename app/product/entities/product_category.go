package entities

type ProductCategory struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
