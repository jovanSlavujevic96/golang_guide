package middlewares

import (
	"net/http"

	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(ginContext *gin.Context) {
	token := ginContext.GetHeader("Authorization")
	if token == "" {
		ginContext.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		ginContext.Abort()
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		ginContext.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		ginContext.Abort()
		return
	}

	ginContext.Set("userId", userId)
	ginContext.Next()
}
