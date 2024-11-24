package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/kechdarho/FinTrack/auth/internal/models"
	"net/http"
)

func (h *Handlers) Register(c *gin.Context) {
	var request models.RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
	}

}
