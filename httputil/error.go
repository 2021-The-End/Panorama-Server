package main

import "github.com/gin-gonic/gin"

type HTTPError struct {
	Code    int    `json:"statuscode" example:"400"`
	Message string `json:"error" example:"bad request"`
}

// NewError example
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}
