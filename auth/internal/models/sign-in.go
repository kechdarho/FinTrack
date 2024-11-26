package models

type SignInRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserSession struct {
	UserID       int    `json:"userID"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
