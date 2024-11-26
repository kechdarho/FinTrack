package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/kechdarho/FinTrack/auth/internal/models"
	"golang.org/x/crypto/bcrypt"
)

const (
	refreshTokenCacheKey = "refresh-token"
	accessTokenCacheKey  = "access-token"
)

func (svc *AuthenticationService) SignIn(ctx context.Context, username, password string) (response models.UserSession, err error) {
	user, err := svc.authPg.GetUser(ctx, username)
	if err != nil {
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		err = errors.New("password isn't correct")
		return
	}

	if _, ok := svc.memoryCache.Get(fmt.Sprintf(accessTokenCacheKey+"%d", user.UserID)); ok {

	}

	accessToken, err := svc.jwtWrapper.GenerateAccessToken(user.UserID)
	if err != nil {
		return
	}

	svc.memoryCache.Set(fmt.Sprintf(accessTokenCacheKey+"%d", user.UserID), accessToken, accessToken.TTL)

	refreshToken, err := svc.jwtWrapper.GenerateRefreshToken(user.UserID)
	if err != nil {
		return
	}

	svc.memoryCache.Set(fmt.Sprintf(refreshTokenCacheKey+"%d", user.UserID), refreshToken, refreshToken.TTL)

	response = models.UserSession{
		UserID:       user.UserID,
		RefreshToken: refreshToken.Token,
		AccessToken:  accessToken.Token,
	}

	return
}
