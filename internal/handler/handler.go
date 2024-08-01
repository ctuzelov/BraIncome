package handler

import (
	"braincome/internal/service"
	"log"
	"math"
	"text/template"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	ErrorLog  *log.Logger
	services  *service.Service
	Tempcache *template.Template
}

func Ceil(x float64) float64 {
	return math.Ceil(x)
}

func Mod(x, y int) int {
	return x % y
}

func NewHandler(logger *log.Logger, services *service.Service) (*Handler, error) {
	funcMap := template.FuncMap{
		"seq": func(start, end int) []int {
			s := make([]int, end-start+1)
			for i := range s {
				s[i] = start + i
			}
			return s
		},
		"ceil": Ceil,
		"mod":  Mod,
	}

	tempcache, err := template.New("").Funcs(funcMap).ParseGlob("assets/html/*.html")
	return &Handler{logger, services, tempcache}, err
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Static("/assets", "assets/")

	router.Use(h.Middleware)

	router.GET("/", h.HomePage)

	router.GET("/contact", h.ContactPage)
	router.POST("/contact", h.ContactFormHandler)

	router.GET("/courses", h.Courses)
	router.GET("/courses/:id", h.Course)

	router.GET("/course/publish", h.IsAdminMiddleware, h.PublishPage)
	router.POST("/course/publish", h.IsAdminMiddleware, h.CourseFormHandler)

	router.GET("/sign-up", h.SignUpPage)
	router.GET("/sign-in", h.SignInPage)
	router.POST("/sign-up", h.SignUp)
	router.POST("/sign-in", h.SignIn)

	router.GET("/sign-out", h.SignOut)

	router.GET("/user/:id", h.MyProfile)
	router.GET("/user/make-instructor", h.IsAdminMiddleware, h.AddInstructorPage)
	router.POST("/user/make-instructor", h.IsAdminMiddleware, h.AddInstructor)
	router.GET("/user/grant-admin-privileges", h.GrantAdminPrivileges)
	return router
}
