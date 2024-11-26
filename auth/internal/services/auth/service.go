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
	CreateUser(ctx context.Context, email, username, password string) (userID int, err error)
	GetUser(ctx context.Context, login string) (user models.User, err error)
}

type memoryCache interface {
	Get(k string) (interface{}, bool)
	Set(k string, x interface{}, d time.Duration)
	Delete(k string)
}

type WrapperJWT interface {
	GenerateAccessToken(userID int) (accessToken models.AccessToken, err error)
	GenerateRefreshToken(userID int) (refreshToken models.RefreshToken, err error)
	ValidateAccessToken(tokenString string) (int, error)
	ValidateRefreshToken(tokenString string) (int, error)
}

type AuthenticationSrv interface {
	SignUp(ctx context.Context, registerRequest models.SignUpRequest) (response models.SignUpResponse, err error)
	SignIn(ctx context.Context, username, password string) (response models.UserSession, err error)
}
