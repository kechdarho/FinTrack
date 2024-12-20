package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	BlackListRefreshTokenCacheKey = "black-list-refresh-token"
	BlackListAccessTokenCacheKey  = "black-list-access-token"
)

type JwtService interface {
	GenerateAccessToken(userID int64) (string, error)
	GenerateRefreshToken(userID int64) (string, error)
	ValidateAccessToken(tokenString string) (int64, error)
	ValidateRefreshToken(tokenString string) (int64, error)
}

type WrapperJWT struct {
	SecretKey           string
	AccessTokenExpires  time.Duration
	RefreshTokenExpires time.Duration
}

func NewJWTWrapper(secretKey string, accessTokenExpires, refreshTokenExpires time.Duration) *WrapperJWT {
	return &WrapperJWT{
		SecretKey:           secretKey,
		AccessTokenExpires:  accessTokenExpires,
		RefreshTokenExpires: refreshTokenExpires,
	}
}

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func (j *WrapperJWT) GenerateAccessToken(userID uint) (accessToken string, err error) {
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.AccessTokenExpires * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.SecretKey))
}

func (j *WrapperJWT) GenerateRefreshToken(userID uint) (refreshToken string, err error) {
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.RefreshTokenExpires * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.SecretKey))
}

func (j *WrapperJWT) ValidateAccessToken(tokenString string) (uint, error) {
	return j.validateToken(tokenString)
}

func (j *WrapperJWT) ValidateRefreshToken(tokenString string) (uint, error) {
	return j.validateToken(tokenString)
}

func (j *WrapperJWT) validateToken(tokenString string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.UserID, nil
	}

	return 0, err
}
