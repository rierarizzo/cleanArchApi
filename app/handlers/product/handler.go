package product

import "github.com/labstack/echo/v4"

type ProductHandler interface {
	GetProducts(c echo.Context) error
}
