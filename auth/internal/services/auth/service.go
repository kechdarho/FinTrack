package auth

import (
	"context"
	"github.com/kechdarho/FinTrack/auth/internal/models"
)

type authService struct {
	memoryCache memoryCache
	authPg      authPg
}

func NewAuthService(memoryCache memoryCache, authPg authPg) AuthenticationService {
	return &authService{
		memoryCache: memoryCache,
		authPg:      authPg,
	}
}

type authPg interface {
	CreateUser(username, password, phone, email string) (success bool, err error)
}

type memoryCache interface {
}

type AuthenticationService interface {
	Login(username, password string) (jwt string, err error)
	Logout(token string) (err error)
	Register(ctx context.Context, registerRequest models.RegisterRequest) (result models.RegisterResponse, err error)
}
