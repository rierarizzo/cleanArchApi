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

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product productDomain.Product

	if err := decoder.Bind(w, r, &product); err != nil {
		responder.ErrorJSON(w, errorDomain.ErrBadRequest, http.StatusBadRequest)
	}

	err := Usecase.CreateProduct(&product)
	if err != nil {
		responder.ErrorJSON(w, errorDomain.ErrBadRequest, http.StatusBadRequest)
	}

	responder.WriteJSON(w, product, http.StatusOK)
}

func CreateProductCategory(w http.ResponseWriter, r *http.Request) {
	var category productDomain.Category

	if err := decoder.Bind(w, r, &category); err != nil {
		responder.ErrorJSON(w, errorDomain.ErrBadRequest, http.StatusBadRequest)
	}

	err := Usecase.CreateProductCategory(&category)
	if err != nil {
		responder.ErrorJSON(w, errorDomain.ErrBadRequest, http.StatusBadRequest)
	}

	responder.WriteJSON(w, category, http.StatusOK)
}

func CreateProductSubcategory(w http.ResponseWriter, r *http.Request) {
	var subcategory productDomain.Subcategory

	if err := decoder.Bind(w, r, &subcategory); err != nil {
		responder.ErrorJSON(w, errorDomain.ErrBadRequest, http.StatusBadRequest)
	}

	err := Usecase.CreateProductSubcategory(&subcategory)
	if err != nil {
		responder.ErrorJSON(w, errorDomain.ErrBadRequest, http.StatusBadRequest)
	}

	responder.WriteJSON(w, subcategory, http.StatusOK)
}

func CreateProductSource(w http.ResponseWriter, r *http.Request) {
	var source productDomain.Source

	if err := decoder.Bind(w, r, &source); err != nil {
		responder.ErrorJSON(w, errorDomain.ErrBadRequest, http.StatusBadRequest)
	}

	err := Usecase.CreateProductSource(&source)
	if err != nil {
		responder.ErrorJSON(w, errorDomain.ErrBadRequest, http.StatusBadRequest)
	}

	responder.WriteJSON(w, source, http.StatusOK)
}
