package http

import (
	"log/slog"
	errorDomain "myclothing/api/domain/error"
	productDomain "myclothing/api/domain/product"
	productHandler "myclothing/api/server/handlers/product"
	"myclothing/api/server/helpers/responder"
	productUsecase "myclothing/api/usecases/product"
	"net/http"
)

type productHttpHandler struct {
	productUsecase productUsecase.Usecase
}

func NewProductHttpHandler(productUsecase productUsecase.Usecase) productHandler.Handler {
	return &productHttpHandler{productUsecase: productUsecase}
}

func (h *productHttpHandler) GetProducts(w http.ResponseWriter, _ *http.Request) {
	products, err := h.productUsecase.GetProducts()
	if err != nil {
		if err = responder.ErrorJSON(w, err, http.StatusInternalServerError); err != nil {
			slog.Error("Error while writing error JSON:", err)
		}
	}

	if err = responder.WriteJSON(w, products, http.StatusOK); err != nil {
		slog.Error("Error while writing JSON:", err)
	}
}

func (h *productHttpHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product productDomain.Product
	if err := responder.Bind(w, r, &product); err != nil {
		if err = responder.ErrorJSON(w, errorDomain.ErrBadRequest, http.StatusBadRequest); err != nil {
			slog.Error("Error while writing error JSON:", err)
		}
	}

	err := h.productUsecase.CreateProduct(&product)
	if err != nil {
		if err = responder.ErrorJSON(w, errorDomain.ErrBadRequest, http.StatusBadRequest); err != nil {
			slog.Error("Error while writing error JSON:", err)
		}
	}

	if err = responder.WriteJSON(w, product, http.StatusOK); err != nil {
		slog.Error("Error while writing JSON:", err)
	}
}

func (h *productHttpHandler) CreateProductCategory(w http.ResponseWriter, r *http.Request) {
	var category productDomain.Category
	if err := responder.Bind(w, r, &category); err != nil {
		if err = responder.ErrorJSON(w, errorDomain.ErrBadRequest, http.StatusBadRequest); err != nil {
			slog.Error("Error while writing error JSON:", err)
		}
	}

	err := h.productUsecase.CreateProductCategory(&category)
	if err != nil {
		if err = responder.ErrorJSON(w, errorDomain.ErrBadRequest, http.StatusBadRequest); err != nil {
			slog.Error("Error while writing error JSON:", err)
		}
	}

	if err = responder.WriteJSON(w, category, http.StatusOK); err != nil {
		slog.Error("Error while writing JSON:", err)
	}
}

func (h *productHttpHandler) CreateProductSubcategory(w http.ResponseWriter, r *http.Request) {
	var subcategory productDomain.Subcategory
	if err := responder.Bind(w, r, &subcategory); err != nil {
		if err = responder.ErrorJSON(w, errorDomain.ErrBadRequest, http.StatusBadRequest); err != nil {
			slog.Error("Error while writing error JSON:", err)
		}
	}

	err := h.productUsecase.CreateProductSubcategory(&subcategory)
	if err != nil {
		if err = responder.ErrorJSON(w, errorDomain.ErrBadRequest, http.StatusBadRequest); err != nil {
			slog.Error("Error while writing error JSON:", err)
		}
	}

	if err = responder.WriteJSON(w, subcategory, http.StatusOK); err != nil {
		slog.Error("Error while writing JSON:", err)
	}
}

func (h *productHttpHandler) CreateProductSource(w http.ResponseWriter, r *http.Request) {
	var source productDomain.Source
	if err := responder.Bind(w, r, &source); err != nil {
		if err = responder.ErrorJSON(w, errorDomain.ErrBadRequest, http.StatusBadRequest); err != nil {
			slog.Error("Error while writing error JSON:", err)
		}
	}

	err := h.productUsecase.CreateProductSource(&source)
	if err != nil {
		if err = responder.ErrorJSON(w, errorDomain.ErrBadRequest, http.StatusBadRequest); err != nil {
			slog.Error("Error while writing error JSON:", err)
		}
	}

	if err = responder.WriteJSON(w, source, http.StatusOK); err != nil {
		slog.Error("Error while writing JSON:", err)
	}
}
