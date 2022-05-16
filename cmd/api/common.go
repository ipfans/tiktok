package api

import (
	"github.com/gin-gonic/gin"
)

type Client struct{}

func NewClient() *Client {
	return &Client{
		// TODO
	}
}

// NewError example
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
