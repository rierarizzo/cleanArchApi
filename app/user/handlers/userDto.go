package handlers

type (
	userResponse struct {
		Id       uint32 `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}
)
