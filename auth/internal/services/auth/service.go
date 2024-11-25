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
	CreateUser(ctx context.Context, email, username, password string) (userID int, err error)
	GetUser(ctx context.Context, login string) (user models.User, err error)
}

type memoryCache interface {
}

type AuthenticationService interface {
	SignUp(ctx context.Context, registerRequest models.SignUpRequest) (response models.SignUpResponse, err error)
	SignIn(ctx context.Context, username, password string) (response models.SignInResponse, err error)
}
