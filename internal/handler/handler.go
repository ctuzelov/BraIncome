package handler

import (
	"braincome/internal/service"
	"log"
	"text/template"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	ErrorLog  *log.Logger
	services  *service.Service
	Tempcache *template.Template
}

func NewHandler(logger *log.Logger, services *service.Service) (*Handler, error) {
	tempcache, err := template.ParseGlob("assets/html/*.html")
	return &Handler{logger, services, tempcache}, err
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Static("/assets", "assets/")

	router.GET("/", h.HomePage)
	router.GET("/courses", h.Courses)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	// Define a new group for admin-only routes
	admin := router.Group("/admin")
	admin.Use(h.Authenticate) // Apply the isAdmin middleware to the admin group

	// // Add the admin-only routes here
	admin.GET("/user/:user_id", h.GetUser)
	// admin.POST("/poste/video", h.PostVideo)
	// admin.GET("/users", h.GetUsers)

	return router
}
