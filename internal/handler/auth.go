package handler

import (
	"braincome/internal/models"
	"braincome/internal/validator"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUpPage(c *gin.Context) {
	data := c.MustGet("data").(*Data)
	h.TemplateRender(c, http.StatusOK, "sign-up.html", data)
}

func (h *Handler) SignInPage(c *gin.Context) {
	data := c.MustGet("data").(*Data)
	h.TemplateRender(c, http.StatusOK, "log-in.html", data)
}

func (h *Handler) SignUp(c *gin.Context) {
	rawJSON, _ := c.GetRawData()
	fmt.Println(string(rawJSON))
	var form models.User

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(form.Email, form.First_name, form.Last_name, form.Password)

	data := c.MustGet("data").(*Data)
	data.Content = form

	if form.Email == "" || form.First_name == "" || form.Last_name == "" || form.Password == "" {
		h.errorpage(c, http.StatusBadRequest, nil)
		return
	}

	data.ErrMsgs = validator.GetErrMsgs(form)
	if len(data.ErrMsgs) != 0 {
		h.TemplateRender(c, http.StatusUnprocessableEntity, "sign-up.html", data)
		return
	}

	if err := h.services.SignUp(form); err != nil {
		switch err {
		case models.ErrDuplicateEmail:
			data.ErrMsgs["email"] = validator.MsgEmailExists
		case models.ErrDuplicateName:
			data.ErrMsgs["name"] = validator.MsgNameExists
		default:
			c.Error(err)
			h.errorpage(c, http.StatusInternalServerError, err)
			return
		}
		h.TemplateRender(c, http.StatusConflict, "sign-up.html", data)
		return
	}

	c.Redirect(http.StatusSeeOther, "/sign-in")
}

func (h *Handler) SignIn(c *gin.Context) {
	var form models.User

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := c.MustGet("data").(*Data)
	fmt.Println(form)
	data.Content = form

	if form.Email == "" || form.Password == "" {
		h.errorpage(c, http.StatusBadRequest, nil)
		return
	}

	user, err := h.services.SignIn(form.Email, form.Password)
	if err != nil {
		data.ErrMsgs = make(map[string]string)
		switch err {
		case models.ErrNoRecord:
			data.ErrMsgs["login"] = validator.MsgUserNotFound
		case models.ErrInvalidCredentials:
			data.ErrMsgs["password"] = validator.MsgNotCorrectPassword
		default:
			c.Error(err)
			h.errorpage(c, http.StatusInternalServerError, err)
			return
		}
		h.TemplateRender(c, http.StatusUnauthorized, "sign-in.html", data)
		return
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: *user.Token,
		Path:  "/",
	}
	c.SetCookie(cookie.Name, cookie.Value, 0, cookie.Path, cookie.Domain, false, false)

	c.Redirect(http.StatusSeeOther, "/")
}

func (h *Handler) SignOut(c *gin.Context) {
	data := c.MustGet("data").(*Data)

	if data.User.Email == "" && data.User.First_name == "" && len(data.User.AccessibleVideos) == 0 {
		h.errorpage(c, http.StatusUnauthorized, nil)
		return
	}

	err := h.services.LogOut(*data.User.Token)
	if err != nil {
		c.Error(err)
		h.errorpage(c, http.StatusInternalServerError, err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}
