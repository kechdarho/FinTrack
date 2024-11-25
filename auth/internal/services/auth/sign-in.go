package auth

import (
	"context"
	"errors"
	"github.com/kechdarho/FinTrack/auth/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (svc *AuthService) SignIn(ctx context.Context, username, password string) (response models.SignInResponse, err error) {
	user, err := svc.authPg.GetUser(ctx, username)
	if err != nil {
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		err = errors.New("password isn't correct")
		return
	}

	return
}
