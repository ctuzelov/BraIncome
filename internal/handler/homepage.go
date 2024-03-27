package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HomePage(c *gin.Context) {
	data := c.MustGet("data").(*Data)


	PopuralCourses, err := h.services.Courses.GetSeveral(4, -1)
	if err != nil{
		h.errorpage(c, http.StatusInternalServerError, err)
		return
	}

	data.Content = PopuralCourses

	h.TemplateRender(c, http.StatusOK, "index.html", data)
}
