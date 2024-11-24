package handlers

import (
	"context"
	"github.com/kechdarho/FinTrack/auth/internal/models"
)

type authService interface {
	Register(ctx context.Context, registerRequest models.RegisterRequest) (result models.RegisterResponse, err error)
}
