package handler

import (
	"braincome/internal/models"
	"braincome/util"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Middleware(c *gin.Context) {
	token, err := c.Cookie("token")
	data := &Data{}

	// fmt.Println("fghrgjrtjrth")

	switch err {
	case http.ErrNoCookie:
		data.User = models.User{}
		data.IsAuthorized = false
		data.IsAdmin = false
	case nil:
		validToken, err := util.ValidateToken(token)
		if err != nil {
			h.errorpage(c, http.StatusBadRequest, err)
		}

		id := validToken["uid"].(string)
		ID, _ := util.StringToObjectID(id)
		data.User.Email, data.User.User_type, data.User.First_name, data.User.ID = validToken["email"].(string), validToken["user_type"].(string), validToken["name"].(string), ID
		data.IsAuthorized = true
		data.IsAdmin = data.User.User_type == "admin"
	}

	c.Set("data", data)

	c.Next()
}

func (h *Handler) IsAdminMiddleware(c *gin.Context) {
	data := c.MustGet("data").(*Data)

	if data.User.Email != "chingizkhan.tuzelov@gmail.com" {
		h.errorpage(c, http.StatusForbidden, errors.New("admin access required"))
		c.Abort()
		return
	}

	c.Next()
}
