package handlers

import (
	"github.com/labstack/echo/v4"
	"myclothing/app/helpers/http/responder"
	"myclothing/app/product/usecases"
	"net/http"
)

type productHttpHandler struct {
	productUsecase usecases.ProductUsecase
}

func NewProductHttpHandler(productUsecase usecases.ProductUsecase) ProductHandler {
	return &productHttpHandler{productUsecase: productUsecase}
}

func (h *productHttpHandler) GetProducts(c echo.Context) error {
	products, err := h.productUsecase.GetProducts()
	if err != nil {
		return err
	}

	return responder.Ok(c, http.StatusOK, products)
}
