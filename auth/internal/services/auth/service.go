package auth

import (
	"context"
	"github.com/kechdarho/FinTrack/auth/internal/models"
	"time"
)

type AuthenticationService struct {
	memoryCache memoryCache
	authPg      authPgStorage
	jwtWrapper  WrapperJWT
}

func NewAuthService(memoryCache memoryCache, authPg authPgStorage, jwtWrapper WrapperJWT) *AuthenticationService {
	return &AuthenticationService{
		memoryCache: memoryCache,
		authPg:      authPg,
		jwtWrapper:  jwtWrapper,
	}
}

type authPgStorage interface {
	CreateUser(ctx context.Context, email, username, password string) (userID uint, err error)
	GetUser(ctx context.Context, login string) (user models.User, err error)
}

type memoryCache interface {
	Get(k string) (interface{}, bool)
	Set(k string, x interface{}, d time.Duration)
	Delete(k string)
}

type WrapperJWT interface {
	GenerateAccessToken(userID uint) (accessToken string, err error)
	GenerateRefreshToken(userID uint) (refreshToken string, err error)
	ValidateAccessToken(tokenString string) (uint, error)
	ValidateRefreshToken(tokenString string) (uint, error)
}

type AuthenticationSrv interface {
	SignUp(ctx context.Context, registerRequest models.SignUpRequest) (response models.SignUpResponse, err error)
	SignIn(ctx context.Context, username, password string) (response models.SignInResponse, err error)
}
