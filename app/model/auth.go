package model

type AuthRequest struct {
	Username string `json:"username" example:"wahyu"`
	Password string `json:"password" example:"strongpassword"`
}

type AuthResponse struct {
	ID       uint
	Username string
	Password string
}
