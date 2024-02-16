package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: add contact form logic

func (h *Handler) ContactPage(c *gin.Context) {
	data := c.MustGet("data").(*Data)
	h.TemplateRender(c, http.StatusOK, "contact.html", data)
}

func (h *Handler) Contact(c *gin.Context) {
	h.TemplateRender(c, http.StatusOK, "contact.html", nil)
}
