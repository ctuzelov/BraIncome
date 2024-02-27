package handler

import (
	"braincome/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO: logic for sorting and fintering courses by category
// TODO: logic for searching courses by name
func (h *Handler) Courses(c *gin.Context) {
	data := c.MustGet("data").(*Data)

	// Check if a category filter is applied
	category := c.Query("category")

	// Retrieve courses based on the category filter
	var courses []models.Course
	var err error
	if category != "" {
		courses, err = h.services.GetByCategory(category)
	} else {
		courses, err = h.services.GetAll()
	}

	if err != nil {
		// Handle the error (e.g., log it, render an error page)
		h.errorpage(c, http.StatusInternalServerError, err)
		return
	}

	// Pass the courses data to the template
	data.Content = courses

	// Render the courses page template with the data
	h.TemplateRender(c, http.StatusOK, "courses.html", data)
}

func (h *Handler) Course(c *gin.Context) {
	data := c.MustGet("data").(*Data)

	courseID := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(courseID)
	// Retrieve the course from the database
	course, err := h.services.Courses.GetById(objectID)

	if err != nil {
		h.errorpage(c, http.StatusNotFound, err)
	}

	data.Content = course
	h.TemplateRender(c, http.StatusOK, "courses-details.html", data)
}
