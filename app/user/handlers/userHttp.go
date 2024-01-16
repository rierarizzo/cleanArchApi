package handlers

import (
	"cleanArchApi/app/helpers/http/responder"
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
	response := make([]userResponse, 0)
	users, err := h.userUsecase.GetAllUsers()

	if err != nil {
		return err
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

	return responder.Ok(c, http.StatusOK, response)
}
