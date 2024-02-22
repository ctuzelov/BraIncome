package handler

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PublishPage(c *gin.Context) {
	data := c.MustGet("data").(*Data)
	h.TemplateRender(c, http.StatusOK, "publish.html", data)
}

type CourseForm struct {
	FullName    string   `form:"fullName"`
	Title       string   `form:"title"`
	Description string   `form:"description"`
	Categories  string   `form:"categories"`
	Language    string   `form:"language"`
	Modules     []Module `form:"modules[]"`
}

type Lesson struct {
	Name string `form:"name"`
	Link string `form:"link"`
}

type Module struct {
	Name    string   `form:"name"`
	Lessons []Lesson `form:"lessons[]"`
}

func (h *Handler) Publish(c *gin.Context) {
	body, _ := c.GetRawData()

	formData, err := url.ParseQuery(string(body))
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	// Create a CourseForm struct and populate it with the decoded form data
	var courseForm CourseForm
	courseForm.FullName = formData.Get("fullName")
	courseForm.Title = formData.Get("title")
	courseForm.Description = formData.Get("description")
	courseForm.Categories = formData.Get("categories")
	courseForm.Language = formData.Get("language")

	for i := 0; ; i++ {
		moduleNameKey := fmt.Sprintf("modules[%d][name]", i)
		moduleName := formData.Get(moduleNameKey)
		if moduleName == "" {
			break
		}

		module := Module{Name: moduleName}

		for j := 0; ; j++ {
			lessonNameKey := fmt.Sprintf("modules[%d][lessons][%d][name]", i, j)
			lessonName := formData.Get(lessonNameKey)
			fmt.Println(lessonName, lessonNameKey)
			if lessonName == "" {
				break
			}

			lessonLinkKey := fmt.Sprintf("modules[%d][lessons][%d][link]", i, j)
			lessonLink := formData.Get(lessonLinkKey)

			module.Lessons = append(module.Lessons, Lesson{Name: lessonName, Link: lessonLink})
		}

		courseForm.Modules = append(courseForm.Modules, module)
	}

	fmt.Println(courseForm.Modules)
	c.JSON(http.StatusOK, gin.H{
		"message": "Form data parsed successfully",
		"form":    courseForm,
	})
}

func (h *Handler) GrantAdminPrivileges(c *gin.Context) {
	cookie, err := c.Request.Cookie("session")
	if err != nil {
		h.errorpage(c, http.StatusInternalServerError, err)
	}
	data := &Data{}

	data.User, err = h.services.User.GetByToken(cookie.Value)

	if err != nil {
		h.errorpage(c, http.StatusInternalServerError, err)
	}

	h.services.MakeAdmin(data.User.Email)

	c.Redirect(http.StatusSeeOther, "/")
}
