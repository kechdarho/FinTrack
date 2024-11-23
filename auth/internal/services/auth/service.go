package auth

import (
	"auth/internal/models"
	"context"
)

type AuthService struct {
	appCache appCache
	authPg   authPg
}

func NewAuthService(authPg authPg) *AuthService {
	return &AuthService{
		authPg: authPg,
	}
}

type authPg interface {
	CreateUser(username, password, phone, email string) (success bool, err error)
}

type AppCache interface {
	Get(stri)
	set()
}

type AuthenticationService interface {
	Login(username, password string) (jwt string, err error)
	Logout(token string) (err error)
	Register(ctx context.Context, registerRequest models.RegisterRequest) (result models.RegisterResponse, err error)
}
