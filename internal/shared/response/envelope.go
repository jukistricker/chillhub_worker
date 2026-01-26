package response

import "github.com/gin-gonic/gin"

type Envelope struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func Send(c *gin.Context, status int, msg string, data any) {
	c.JSON(status, Envelope{
		Status:  status,
		Message: msg,
		Data:    data,
	})
}