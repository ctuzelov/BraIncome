package handler

import (
	"braincome/internal/models"
	"braincome/internal/validator"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validator.GetErrMsgs(user)
	if len(validationErr) != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
		return
	}

	resultInsertionNumber, err := h.services.CreateUser(user)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resultInsertionNumber)
}

func (h *Handler) SignIn(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	foundUser, err := h.services.GetUserByEmail(user.Email, user.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, foundUser)
}
