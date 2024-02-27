package handler

import (
	"braincome/internal/models"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PublishPage(c *gin.Context) {
	data := c.MustGet("data").(*Data)
	h.TemplateRender(c, http.StatusOK, "publish.html", data)
}

type CourseForm struct {
	FullName       string          `form:"fullName"`
	Email          string          `form:"email"`
	Title          string          `form:"title"`
	Description    string          `form:"description"`
	Categories     []string        `form:"categories"`
	Language       string          `form:"language"`
	CoverPhotoLink string          `form:"coverPhotoLink"`
	Modules        []models.Module `form:"modules[]"`
}

func (h *Handler) CourseFormHandler(c *gin.Context) {
	body, _ := c.GetRawData()

	formData, err := url.ParseQuery(string(body))
	if err != nil {
		h.errorpage(c, http.StatusBadRequest, err)
		return
	}

	// Create a CourseForm struct and populate it with the decoded form data
	var courseForm CourseForm
	courseForm.FullName = formData.Get("fullName")
	courseForm.Email = formData.Get("email")
	courseForm.Title = formData.Get("title")
	courseForm.Description = formData.Get("description")
	courseForm.Categories = strings.Split(formData.Get("categories"), " ")
	courseForm.Language = formData.Get("language")
	courseForm.CoverPhotoLink = GenerateThumbnailURL(ExtractPhotoID(formData.Get("coverPhotoLink")))

	for i := 0; ; i++ {
		moduleNameKey := fmt.Sprintf("modules[%d][name]", i)
		moduleName := formData.Get(moduleNameKey)
		if moduleName == "" {
			break
		}

		module := models.Module{Name: moduleName}

		for j := 0; ; j++ {
			lessonNameKey := fmt.Sprintf("modules[%d][lessons][%d][name]", i, j)
			lessonName := formData.Get(lessonNameKey)
			fmt.Println(lessonName, lessonNameKey)
			if lessonName == "" {
				break
			}

			lessonLinkKey := fmt.Sprintf("modules[%d][lessons][%d][link]", i, j)
			lessonLink := formData.Get(lessonLinkKey)

			module.Lessons = append(module.Lessons, models.Lesson{Name: lessonName, Link: lessonLink})
		}

		courseForm.Modules = append(courseForm.Modules, module)
	}

	instructor, _ := h.services.Instructor.GetByEmail(courseForm.Email)

	proto := models.Course{
		Instructor:     instructor,
		Title:          courseForm.Title,
		Description:    courseForm.Description,
		Categories:     courseForm.Categories,
		Language:       courseForm.Language,
		CoverPhotoLink: courseForm.CoverPhotoLink,
		Curriculum:     courseForm.Modules,
	}

	err = h.services.Courses.Insert(proto)

	if err != nil {
		h.errorpage(c, http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{"data": proto})
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
