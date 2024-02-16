package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Publish(c *gin.Context) {
	h.TemplateRender(c, http.StatusOK, "publish.html", nil)
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
