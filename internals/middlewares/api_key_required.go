package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiKeyRequired(context *gin.Context) {
	if key := context.Request.Header.Get("X-Api-Key"); key == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		context.Abort()
	}
	context.Next()
}
