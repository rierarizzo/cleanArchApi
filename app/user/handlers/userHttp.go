package handlers

import (
	"cleanArchApi/app/user/usecases"
	"github.com/labstack/echo/v4"
	"net/http"
)

type userHttpHandler struct {
	userUsecase usecases.UserUsecase
}

func NewUserHttpHandler(userUsecase usecases.UserUsecase) UserHandler {
	return &userHttpHandler{userUsecase: userUsecase}
}

func (h *userHttpHandler) GetUsers(c echo.Context) error {
	var response []userResponse
	users, err := h.userUsecase.GetAllUsers()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	for _, user := range users {
		response = append(
			response, userResponse{
				Id:       user.Id,
				Username: user.Username,
				Email:    user.Email,
			},
		)
	}

	return c.JSON(http.StatusOK, response)
}
