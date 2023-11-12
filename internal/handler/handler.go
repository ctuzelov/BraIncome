package handler

import (
	"braincome/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	// Define a new group for admin-only routes
	admin := router.Group("/admin")
	admin.Use(h.Authenticate) // Apply the isAdmin middleware to the admin group

	// // Add the admin-only routes here
	// admin.GET("/users", h.GetUsers)
	// admin.GET("/user/:user_id", h.GetUser)

	return router
}
