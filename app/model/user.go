package model

type UserResponse struct {
	BasicData
	Username string `json:"username"`
}

type UserRequest struct {
	Username string `json:"username" example:"wahyu"`
	Password string `json:"password" example:"strongpassword"`
}
