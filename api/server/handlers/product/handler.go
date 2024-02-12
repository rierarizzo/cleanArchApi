package product

import (
	errorDomain "myclothing/api/domain/error"
	productDomain "myclothing/api/domain/product"
	"myclothing/api/server/helpers/decoder"
	"myclothing/api/server/helpers/responder"
	productUsecase "myclothing/api/usecases/product"
	"net/http"
)

var Usecase productUsecase.Usecase

func GetProducts(w http.ResponseWriter, _ *http.Request) {
	products, err := Usecase.GetProducts()
	if err != nil {
		responder.ErrorJSON(w, err, http.StatusInternalServerError)
	}

	responder.WriteJSON(w, products, http.StatusOK)
}

func createProductElement[T any](w http.ResponseWriter, r *http.Request, element T, usecaseFunc func(element *T) error) {
	if err := decoder.Bind(w, r, &element); err != nil {
		responder.ErrorJSON(w, errorDomain.ErrBadRequest, http.StatusBadRequest)
	}

	err := usecaseFunc(&element)
	if err != nil {
		responder.ErrorJSON(w, errorDomain.ErrBadRequest, http.StatusBadRequest)
	}

	responder.WriteJSON(w, element, http.StatusOK)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product productDomain.Product
	createProductElement(w, r, product, Usecase.CreateProduct)
}

func CreateProductCategory(w http.ResponseWriter, r *http.Request) {
	var category productDomain.Category
	createProductElement(w, r, category, Usecase.CreateProductCategory)
}

func CreateProductSubcategory(w http.ResponseWriter, r *http.Request) {
	var subcategory productDomain.Subcategory
	createProductElement(w, r, subcategory, Usecase.CreateProductSubcategory)
}

func CreateProductSource(w http.ResponseWriter, r *http.Request) {
	var source productDomain.Source
	createProductElement(w, r, source, Usecase.CreateProductSource)
}
