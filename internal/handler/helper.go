package handler

import (
	"braincome/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Data struct {
	User         models.User
	Content      any
	IsAuthorized bool
	IsEmpty      bool
	IsAdmin      bool
	ErrMsgs      map[string]string
	Categories   []*string
}

type ErrorData struct {
	Status  int
	Message string
}

func (h *Handler) errorpage(c *gin.Context, status int, err error) {
	if err != nil {
		h.ErrorLog.Printf("server error: %v", err)
	}

	msg := http.StatusText(status)
	errdata := ErrorData{Status: status, Message: msg}

	h.TemplateRender(c, status, "errors.html", errdata)
}
