package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rest-api-event/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
		return
	}

	userid, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not Authorized"})
		return
	}
	context.Set("userid", userid)
	context.Next()
}
