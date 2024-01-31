package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ContactPage(c *gin.Context) {
	h.TemplateRender(c, http.StatusOK, "contact.html", nil)
}

func (h *Handler) Contact(c *gin.Context) {
	h.TemplateRender(c, http.StatusOK, "contact.html", nil)
}
