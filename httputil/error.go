package httputil

import (
	"github.com/gin-gonic/gin"
)

type HTTPError struct {
	Code    int    `json:"statuscode"`
	Message string `json:"msg"`
}

// NewError example
func NewError(c *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	c.JSON(status, er)
}

func NewRedirect(c *gin.Context, status int, locate string) {
	c.Redirect(status, locate)
}
