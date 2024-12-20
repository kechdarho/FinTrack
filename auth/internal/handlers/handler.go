package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/kechdarho/FinTrack/auth/internal/models"
)

type Handlers struct {
	Router      *gin.Engine
	authService authService
}

func NewHandlers(authService authService) *Handlers {
	router := gin.Default()

	h := &Handlers{
		Router:      router,
		authService: authService,
	}

	h.registerRoutes()

	return h
}

type authService interface {
	SignUp(ctx context.Context, registerRequest models.SignUpRequest) (response models.SignUpResponse, err error)
	SignIn(ctx context.Context, username, password string) (response models.SignInResponse, refreshToken string, err error)
	Logout(ctx context.Context, token uint) (err error)
}
