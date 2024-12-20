package handlers

func (h *Handlers) registerRoutes() {
	h.Router.Group("/api").
		Group("/v1").
		POST("/register", h.SignUp).
		POST("/login", h.SignIn).
		DELETE("/logout", h.Logout)
}
