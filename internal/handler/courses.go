package handler

import (
	"braincome/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: logic for sorting and fintering courses by category
// TODO: logic for searching courses by name
func (h *Handler) Courses(c *gin.Context) {
	data := c.MustGet("data").(*Data)
	h.TemplateRender(c, http.StatusOK, "courses.html", data)
}

func (h *Handler) Course(c *gin.Context) {
	data := c.MustGet("data").(*Data)
	data.Content = &models.Course{}
	h.TemplateRender(c, http.StatusOK, "courses-details.html", data)
}
