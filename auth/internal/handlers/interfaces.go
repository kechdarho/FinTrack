package handlers

import (
	"auth/internal/models"
	"context"
)

type authService interface {
	Register(ctx context.Context, registerRequest models.RegisterRequest) (result models.RegisterResponse, err error)
}
