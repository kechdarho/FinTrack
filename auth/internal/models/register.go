package models

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type RegisterResponse struct {
	Success bool `json:"success"`
	UserID  int  `json:"userID"`
}
