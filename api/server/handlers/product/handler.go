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
	productUsecase productUsecase.Usecase
}

func NewProductHttpHandler(productUsecase productUsecase.Usecase) Handler {
	return &productHttpHandler{productUsecase: productUsecase}
}

func (h *productHttpHandler) GetProducts(w http.ResponseWriter, _ *http.Request) {
	products, err := h.productUsecase.GetProducts()
	if err != nil {
		responder.ErrorJSON(w, err, http.StatusInternalServerError)
	}

	responder.WriteJSON(w, products, http.StatusOK)
}

func (h *productHttpHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product productDomain.Product

	if err := decoder.Bind(w, r, &product); err != nil {
		responder.ErrorJSON(w, errorDomain.ErrBadRequest, http.StatusBadRequest)
	}

	err := h.productUsecase.CreateProduct(&product)
	if err != nil {
		responder.ErrorJSON(w, errorDomain.ErrBadRequest, http.StatusBadRequest)
	}

	responder.WriteJSON(w, product, http.StatusOK)
}

func (h *productHttpHandler) CreateProductCategory(w http.ResponseWriter, r *http.Request) {
	var category productDomain.Category

	if err := decoder.Bind(w, r, &category); err != nil {
		responder.ErrorJSON(w, errorDomain.ErrBadRequest, http.StatusBadRequest)
	}

	err := h.productUsecase.CreateProductCategory(&category)
	if err != nil {
		responder.ErrorJSON(w, errorDomain.ErrBadRequest, http.StatusBadRequest)
	}

	responder.WriteJSON(w, category, http.StatusOK)
}

func (h *productHttpHandler) CreateProductSubcategory(w http.ResponseWriter, r *http.Request) {
	var subcategory productDomain.Subcategory

	if err := decoder.Bind(w, r, &subcategory); err != nil {
		responder.ErrorJSON(w, errorDomain.ErrBadRequest, http.StatusBadRequest)
	}

	err := h.productUsecase.CreateProductSubcategory(&subcategory)
	if err != nil {
		responder.ErrorJSON(w, errorDomain.ErrBadRequest, http.StatusBadRequest)
	}

	responder.WriteJSON(w, subcategory, http.StatusOK)
}

func (h *productHttpHandler) CreateProductSource(w http.ResponseWriter, r *http.Request) {
	var source productDomain.Source

	if err := decoder.Bind(w, r, &source); err != nil {
		responder.ErrorJSON(w, errorDomain.ErrBadRequest, http.StatusBadRequest)
	}

	err := h.productUsecase.CreateProductSource(&source)
	if err != nil {
		responder.ErrorJSON(w, errorDomain.ErrBadRequest, http.StatusBadRequest)
	}

	responder.WriteJSON(w, source, http.StatusOK)
}
