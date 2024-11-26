package auth

import (
	"context"
	"github.com/kechdarho/FinTrack/auth/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (svc *AuthenticationService) SignUp(ctx context.Context, registerRequest models.SignUpRequest) (response models.SignUpResponse, err error) {
	hashedPassword, err := hashPassword(registerRequest.Password)
	if err != nil {
		return
	}
	userID, err := svc.authPg.CreateUser(ctx, registerRequest.Email, registerRequest.Username, string(hashedPassword))
	if err != nil {
		return
	}

	response = models.SignUpResponse{UserID: userID}

	return
}

func hashPassword(password string) (hashedPassword []byte, err error) {
	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	return
}
