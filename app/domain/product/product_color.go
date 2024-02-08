package product

type ProductColor struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Hex  string `json:"hex,omitempty"`
}
