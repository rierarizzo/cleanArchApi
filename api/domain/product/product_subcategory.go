package product

type ProductSubcategory struct {
	Id             int              `json:"id,omitempty"`
	ParentCategory *ProductCategory `json:"parent_category,omitempty"`
	Name           string           `json:"name,omitempty"`
	Description    string           `json:"description,omitempty"`
}
