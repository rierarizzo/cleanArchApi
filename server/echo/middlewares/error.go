package middlewares

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"log/slog"
	appError "myclothing/app/error"
	"myclothing/app/helpers/http/responder"
	"net/http"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := getHttpCodeStatusFromAppError(err)

	responderErr := responder.Error(c, code, err)

	if responderErr != nil {
		slog.Error(fmt.Sprintf("middlewares.customHTTPErrorHandler: %v", responderErr))
	}
}

func getHttpCodeStatusFromAppError(err error) int {
	switch {
	case errors.Is(err, appError.ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, appError.ErrUnknown):
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
