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

	err := h.productUsecase.CreateProduct(&product)
	if err != nil {
		return responder.Error(c, http.StatusInternalServerError, err)
	}

	return responder.Ok(c, http.StatusCreated, product)
}

func (h *productHttpHandler) CreateProductCategory(c echo.Context) error {
	var category productDomain.Category
	if err := c.Bind(&category); err != nil {
		return responder.Error(c, http.StatusBadRequest, errorDomain.ErrBadRequest)
	}

	err := h.productUsecase.CreateProductCategory(&category)
	if err != nil {
		return responder.Error(c, http.StatusInternalServerError, err)
	}

	return responder.Ok(c, http.StatusCreated, category)
}

func (h *productHttpHandler) CreateProductSubcategory(c echo.Context) error {
	var subcategory productDomain.Subcategory
	if err := c.Bind(&subcategory); err != nil {
		return responder.Error(c, http.StatusBadRequest, errorDomain.ErrBadRequest)
	}

	err := h.productUsecase.CreateProductSubcategory(&subcategory)
	if err != nil {
		return responder.Error(c, http.StatusInternalServerError, err)
	}

	return responder.Ok(c, http.StatusCreated, subcategory)
}

func (h *productHttpHandler) CreateProductSource(c echo.Context) error {
	var source productDomain.Source
	if err := c.Bind(&source); err != nil {
		return responder.Error(c, http.StatusBadRequest, errorDomain.ErrBadRequest)
	}

	err := h.productUsecase.CreateProductSource(&source)
	if err != nil {
		return responder.Error(c, http.StatusInternalServerError, err)
	}

	return responder.Ok(c, http.StatusCreated, source)
}
