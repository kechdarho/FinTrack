package models

type LogoutRequest struct {
	UserID uint `json:"user_id"`
}

type LogoutResponse struct {
	IsSuccess bool `json:"is_success"`
}
