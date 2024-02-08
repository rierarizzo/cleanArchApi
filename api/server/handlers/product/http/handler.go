package http

import (
	"github.com/labstack/echo/v4"
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

func (h *productHttpHandler) GetProducts(c echo.Context) error {
	products, err := h.productUsecase.GetProducts()
	if err != nil {
		return err
	}

	return responder.Ok(c, http.StatusOK, products)
}

func (h *productHttpHandler) CreateProduct(c echo.Context) error {
	var product productDomain.Product
	if err := c.Bind(&product); err != nil {
		return responder.Error(c, http.StatusBadRequest, errorDomain.ErrBadRequest)
	}

	return responder.Ok(c, http.StatusOK, product)
}
