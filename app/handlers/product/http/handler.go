package http

import (
	"github.com/labstack/echo/v4"
	product2 "myclothing/app/handlers/product"
	"myclothing/app/helpers/http/responder"
	"myclothing/app/usecases/product"
	"net/http"
)

type productHttpHandler struct {
	productUsecase product.Usecase
}

func NewProductHttpHandler(productUsecase product.Usecase) product2.ProductHandler {
	return &productHttpHandler{productUsecase: productUsecase}
}

func (h *productHttpHandler) GetProducts(c echo.Context) error {
	products, err := h.productUsecase.GetProducts()
	if err != nil {
		return err
	}

	return responder.Ok(c, http.StatusOK, products)
}
