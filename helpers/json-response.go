package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonResponse struct{}

var Response = JsonResponse{}

func (j *JsonResponse) SuccessResponse(c *gin.Context, data any, message string) {
	if message == "" {
		message = "Successful"
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func (j *JsonResponse) ErrorResponse(c *gin.Context, err error, statusCode ...int) {
	code := http.StatusBadRequest
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	c.JSON(code, gin.H{
		"success": false,
		"message": err.Error(),
	})
}

func (j *JsonResponse) ServerErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"success": false,
		"message": err.Error(),
	})
}
