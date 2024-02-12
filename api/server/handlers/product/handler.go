package product

import "github.com/labstack/echo/v4"

type Handler interface {
	GetProducts(c echo.Context) error
	CreateProduct(c echo.Context) error
	CreateProductCategory(c echo.Context) error
	CreateProductSubcategory(c echo.Context) error
	CreateProductSource(c echo.Context) error
}
