package handlers

import "github.com/labstack/echo/v4"

type UserHandler interface {
	GetUsers(c echo.Context) error
}
