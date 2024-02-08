package http

import (
	"github.com/labstack/echo/v4"
	productHandler "myclothing/api/server/handlers/product"
	"myclothing/api/server/helpers/responder"
	"myclothing/api/usecases/product"
	"net/http"
)

type productHttpHandler struct {
	productUsecase product.Usecase
}

func NewProductHttpHandler(productUsecase product.Usecase) productHandler.Handler {
	return &productHttpHandler{productUsecase: productUsecase}
}

func (h *productHttpHandler) GetProducts(c echo.Context) error {
	products, err := h.productUsecase.GetProducts()
	if err != nil {
		return err
	}

	return responder.Ok(c, http.StatusOK, products)
}
