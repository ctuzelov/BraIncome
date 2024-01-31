package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HomePage(c *gin.Context) {
	data := c.MustGet("data").(*Data)
	h.TemplateRender(c, http.StatusOK, "index.html", data)
}
