package product

import (
	"net/http"
)

type Handler interface {
	GetProducts(w http.ResponseWriter, r *http.Request)
	CreateProduct(w http.ResponseWriter, r *http.Request)
	CreateProductCategory(w http.ResponseWriter, r *http.Request)
	CreateProductSubcategory(w http.ResponseWriter, r *http.Request)
	CreateProductSource(w http.ResponseWriter, r *http.Request)
}
