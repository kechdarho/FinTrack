package models

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type SignUpResponse struct {
	UserID int `json:"userID"`
}
