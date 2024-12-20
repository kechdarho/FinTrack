package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kechdarho/FinTrack/auth/internal/models"
	"net/http"
)

func (h *Handlers) SignUp(c *gin.Context) {
	request, err := validateRegisterRequest(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.authService.SignUp(c.Request.Context(), request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func validateRegisterRequest(c *gin.Context) (request models.SignUpRequest, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return
	}

	switch {
	case request.Email == "":
		return models.SignUpRequest{}, errors.New("email is required")
	case request.Username == "":
		return models.SignUpRequest{}, errors.New("username is required")
	case request.Password == "":
		return models.SignUpRequest{}, errors.New("password is required")
	}

	return
}
