package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kechdarho/FinTrack/auth/internal/models"
	"net/http"
	"strings"
)

func (h *Handlers) SignIn(c *gin.Context) {
	request, err := validateLoginRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := h.authService.SignIn(c.Request.Context(), request.Login, request.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func validateLoginRequest(c *gin.Context) (request models.SignInRequest, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return
	}

	if len(strings.TrimSpace(request.Login)) != len(request.Login) {
		return models.SignInRequest{}, errors.New("login is required and cannot contain only whitespace")
	}

	if len(request.Login) < 6 {
		return models.SignInRequest{}, errors.New("login must be at least 6 characters long")
	}

	if len(strings.TrimSpace(request.Password)) != len(request.Password) {
		return models.SignInRequest{}, errors.New("password is required and cannot contain only whitespace")
	}

	if len(request.Password) < 6 {
		return models.SignInRequest{}, errors.New("password must be at least 6 characters long")
	}

	return
}
