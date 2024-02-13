package product

import (
	errorDomain "myclothing/api/domain/error"
	productDomain "myclothing/api/domain/product"
	"myclothing/api/server/helpers/decoder"
	"myclothing/api/server/helpers/responder"
	productUsecase "myclothing/api/usecases/product"
	"net/http"
)

type productHttpHandler struct {
	usecase productUsecase.Usecase
}

func NewProductHttpHandler(usecase productUsecase.Usecase) Handler {
	return &productHttpHandler{usecase}
}

func (h *productHttpHandler) GetProducts(w http.ResponseWriter, _ *http.Request) {
	products, err := h.usecase.GetProducts()
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

func (h *productHttpHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product productDomain.Product
	createProductElement(w, r, product, h.usecase.CreateProduct)
}

func (h *productHttpHandler) CreateProductCategory(w http.ResponseWriter, r *http.Request) {
	var category productDomain.Category
	createProductElement(w, r, category, h.usecase.CreateProductCategory)
}

func (h *productHttpHandler) CreateProductSubcategory(w http.ResponseWriter, r *http.Request) {
	var subcategory productDomain.Subcategory
	createProductElement(w, r, subcategory, h.usecase.CreateProductSubcategory)
}

func (h *productHttpHandler) CreateProductSource(w http.ResponseWriter, r *http.Request) {
	var source productDomain.Source
	createProductElement(w, r, source, h.usecase.CreateProductSource)
}
