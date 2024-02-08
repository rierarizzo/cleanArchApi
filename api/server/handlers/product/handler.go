package product

import "github.com/labstack/echo/v4"

type Handler interface {
	GetProducts(c echo.Context) error
}
