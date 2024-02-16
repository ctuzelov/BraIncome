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

	router.Use(h.Middleware)

	router.GET("/", h.HomePage)
	router.GET("/courses", h.Courses)
	router.GET("/courses-details", h.Course)

	router.GET("/contact", h.ContactPage)
	router.POST("/contact", h.Contact)

	router.GET("/sign-up", h.SignUpPage)
	router.GET("/sign-in", h.SignInPage)
	router.POST("/sign-up", h.SignUp)
	router.POST("/sign-in", h.SignIn)

	router.GET("/sign-out", h.SignOut)

	router.GET("/my-profile", h.MyProfile).Use(h.RequireAuth)
	router.GET("/publish", h.Publish).Use(h.IsAdminMiddleware)
	router.POST("/publish", h.Publish).Use(h.IsAdminMiddleware)
	router.GET("/grant-admin-privileges", h.GrantAdminPrivileges)

	admin := router.Group("/admin")
	admin.Use(h.RequireAuth)

	// admin.POST("/post/video", h.PostVideo)
	// admin.GET("/users", h.GetUsers)

	return router
}
