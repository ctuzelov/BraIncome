package handler

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) TemplateRender(c *gin.Context, status int, page string, data interface{}) {
	buf := new(bytes.Buffer)
	err := h.Tempcache.ExecuteTemplate(buf, page, data)
	if err != nil {
		h.ErrorLog.Println(err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.Status(status)
	c.Writer.Write(buf.Bytes())
}
