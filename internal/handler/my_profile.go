package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) MyProfile(c *gin.Context) {
	h.TemplateRender(c, http.StatusOK, "my-profile.html", nil)
}
