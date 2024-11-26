package jwt

import (
	"github.com/kechdarho/FinTrack/auth/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func (j *WrapperJWT) GenerateAccessToken(userID int) (accessToken models.AccessToken, err error) {
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.AccessTokenExpires * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return
	}

	accessToken = models.AccessToken{Token: signedToken, TTL: j.AccessTokenExpires}

	return
}

func (j *WrapperJWT) GenerateRefreshToken(userID int) (refreshToken models.RefreshToken, err error) {
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.RefreshTokenExpires * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return
	}

	refreshToken = models.RefreshToken{Token: signedToken, TTL: j.RefreshTokenExpires}
	return
}

func (j *WrapperJWT) ValidateAccessToken(tokenString string) (int, error) {
	return j.validateToken(tokenString)
}

func (j *WrapperJWT) ValidateRefreshToken(tokenString string) (int, error) {
	return j.validateToken(tokenString)
}

func (j *WrapperJWT) validateToken(tokenString string) (int, error) {
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
