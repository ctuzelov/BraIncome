package handler

import (
	"braincome/internal/models"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddInstructor(c *gin.Context) {
	c.Request.ParseForm()

	var userForm models.Instructor
	userForm.First_name = c.Request.PostForm.Get("first_name")
	userForm.Last_name = c.Request.PostForm.Get("last_name")
	userForm.Email = c.Request.PostForm.Get("email")
	userForm.About = c.Request.PostForm.Get("about")
	userForm.Speciality = c.Request.PostForm.Get("speciality")
	userForm.AvatarLink = c.Request.PostForm.Get("avatar_link")

	// TODO: clarify the method and add validation for userForm
	_, err := h.services.Instructor.GetByEmail(userForm.Email)

	if err != nil && !errors.Is(err, models.ErrDuplicateEmail) {
		c.Redirect(http.StatusSeeOther, "/course/publish")
	}

	err = h.services.Instructor.Insert(userForm)

	if err != nil {
		h.errorpage(c, http.StatusInternalServerError, err)
	}

	c.Redirect(http.StatusSeeOther, "/course/publish")
}

func (h *Handler) AddInstructorPage(c *gin.Context) {
	data := c.MustGet("data").(*Data)

	fmt.Println(data, data.IsAuthorized)
	h.TemplateRender(c, http.StatusOK, "add-instructor.html", data)
}
