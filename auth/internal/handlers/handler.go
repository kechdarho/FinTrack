package handlers

import (
	"github.com/gin-gonic/gin"
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

func (h *Handlers) registerRoutes() {
	h.Router.POST("/register", h.Register)
}
