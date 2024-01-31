package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Courses(c *gin.Context) {
	h.TemplateRender(c, http.StatusOK, "courses.html", nil)
}
