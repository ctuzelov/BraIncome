package handler

import (
	"braincome/internal/models"
	"fmt"
	"net/http"
	"regexp"

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

func ExtractPhotoID(url string) string {
	// Regular expression to extract the PhotoID
	re := regexp.MustCompile(`\/d\/([a-zA-Z0-9_-]+)\/`)
	matches := re.FindStringSubmatch(url)
	if len(matches) < 2 {
		return "" // Return empty string if no match found
	}
	return matches[1]
}

func GenerateThumbnailURL(photoID string) string {
	return fmt.Sprintf("https://drive.google.com/thumbnail?id=%s&sz=w1000", photoID)
}
