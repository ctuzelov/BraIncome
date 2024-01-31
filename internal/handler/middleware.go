package handler

import (
	"braincome/internal/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Middleware(c *gin.Context) {
	cookie, err := c.Request.Cookie("session")
	data := &Data{}

	switch err {
	case http.ErrNoCookie:
		data.User = models.User{}
	case nil:
		data.User, err = h.services.User.GetByToken(cookie.Value)
		if err != nil && !errors.Is(err, models.ErrNoRecord) {
			h.errorpage(c, http.StatusInternalServerError, err)
			c.Abort()
			return
		}
		if data.User.Token != nil {
			data.IsAuthorized = true
		}
	default:
		h.errorpage(c, http.StatusInternalServerError, err)
		c.Abort()
		return
	}

	c.Set("data", data)
	c.Next()
}

func (h *Handler) RequireAuth(c *gin.Context) {
	data, ok := c.Get("data")
	if !ok {
		h.errorpage(c, http.StatusInternalServerError, errors.New("context value not found"))
		c.Abort()
		return
	}

	userData, ok := data.(*Data)
	if !ok {
		h.errorpage(c, http.StatusInternalServerError, errors.New("unexpected data type in context"))
		c.Abort()
		return
	}

	if !userData.IsAuthorized {
		c.Redirect(http.StatusSeeOther, "/signin")
		c.Abort()
		return
	}

	c.Next()
}
