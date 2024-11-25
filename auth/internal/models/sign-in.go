package models

type SignInRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInResponse struct {
	UserID      int    `json:"userID"`
	AccessToken string `json:"access_token"`
}
