package handlers

import (
	"auth/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handlers) Register(c *gin.Context) {
	var request models.RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
	}

}
