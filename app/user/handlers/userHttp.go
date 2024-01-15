package handlers

import (
	httpHelper "cleanArchApi/app/helpers/http"
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
		return httpHelper.Error(c, http.StatusInternalServerError, err)
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

	return httpHelper.Ok(c, http.StatusOK, "Users in DB", response)
}
