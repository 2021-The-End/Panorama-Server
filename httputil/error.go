package httputil

import (
	"github.com/gin-gonic/gin"
)

type HTTPError struct {
	Code    int    `json:"statuscode"`
	Message string `json:"msg"`
}

// NewError example
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}
