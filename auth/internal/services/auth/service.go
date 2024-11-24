package auth

import (
	"context"
	"github.com/kechdarho/FinTrack/auth/internal/models"
)

type AuthService struct {
	memoryCache memoryCache
	authPg      authPgStorage
}

func NewAuthService(memoryCache memoryCache, authPg authPgStorage) *AuthService {
	return &AuthService{
		memoryCache: memoryCache,
		authPg:      authPg,
	}
}

type authPgStorage interface {
}

type memoryCache interface {
}

type AuthenticationService interface {
	Login(username, password string) (jwt string, err error)
	Logout(token string) (err error)
	Register(ctx context.Context, registerRequest models.RegisterRequest) (result models.RegisterResponse, err error)
}
