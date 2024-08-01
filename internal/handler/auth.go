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
	c.Request.ParseForm()

	rawJson, _ := c.GetRawData()
	fmt.Println(rawJson, "raw data")

	firstName := c.Request.PostForm.Get("first_name")
	lastName := c.Request.PostForm.Get("last_name")
	email := c.Request.PostForm.Get("email")
	password := c.Request.PostForm.Get("password")
	// Create a User instance with form data
	var form models.User
	form.First_name = firstName
	form.Last_name = lastName
	form.Email = email
	form.Password = password
	fmt.Println(form.Email, form.First_name, form.Last_name, form.Password)

	data := c.MustGet("data").(*Data)
	data.Content = form

	if form.Email == "" || form.First_name == "" || form.Last_name == "" || form.Password == "" {
		h.errorpage(c, http.StatusBadRequest, nil)
		return
	}

	data.ErrMsgs = validator.GetErrMsgs(form)
	if len(data.ErrMsgs) != 0 {
		fmt.Println(data.ErrMsgs)
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

	fmt.Println(data, "aaaaaaaaaaaaaa")

	c.Redirect(http.StatusSeeOther, "/sign-in")
}

func (h *Handler) SignIn(c *gin.Context) {
	c.Request.ParseForm()
	fmt.Println("sdfasfasfdsafds")
	firstName := c.Request.PostForm.Get("first_name")
	lastName := c.Request.PostForm.Get("last_name")
	email := c.Request.PostForm.Get("email")
	password := c.Request.PostForm.Get("password")
	// Create a User instance with form data
	var form models.User
	form.First_name = firstName
	form.Last_name = lastName
	form.Email = email
	form.Password = password

	data := c.MustGet("data").(*Data)
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
			data.ErrMsgs["email"] = validator.MsgUserNotFound
		case models.ErrInvalidCredentials:
			data.ErrMsgs["password"] = validator.MsgNotCorrectPassword
			fmt.Println("wrong password")
		default:
			c.Error(err)
			h.errorpage(c, http.StatusInternalServerError, err)
			return
		}
		h.TemplateRender(c, http.StatusUnauthorized, "log-in.html", data)
		return
	}

	c.SetCookie("token", user.Token, 3600, "/", "", false, true)

	c.Redirect(http.StatusSeeOther, "/")
}

func (h *Handler) SignOut(c *gin.Context) {
	data := c.MustGet("data").(*Data)

	// if data.User.Email == "" && data.User.First_name == "" {
	// 	h.errorpage(c, http.StatusUnauthorized, nil)
	// 	return
	// }

	err := h.services.DeleteTokensByEmail(data.User.Email)
	if err != nil {
		h.errorpage(c, http.StatusBadRequest, err)
		return
	}

	c.SetCookie("token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)

	c.Redirect(http.StatusSeeOther, "/")
}
