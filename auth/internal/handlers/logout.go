package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handlers) Logout(c *gin.Context) {
	ctx := c.Request.Context()

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing access token"})
	}

	accessToken := strings.TrimPrefix(authHeader, "Bearer ")

	logoutResponse := h.authService.Logout(ctx, accessToken)

	c.JSON(http.StatusOK, logoutResponse)
}
