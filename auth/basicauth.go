package auth

import (
	"os"

	"github.com/gin-gonic/gin"
)

func BasicAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Request.Header.Get("username")
		password := ctx.Request.Header.Get("password")

		if username != os.Getenv("ADMIN_USERNAME") || password != os.Getenv("ADMIN_PASSWORD") {
			if username == "" || password == "" {
				ctx.JSON(401, gin.H{
					"message": "Unauthorized",
				})
				ctx.Abort()
				return
			} else {
				ctx.JSON(401, gin.H{
					"message": "Unauthorized",
				})
				ctx.Abort()
				return
			}
		} else {
			ctx.Header("username", "admin")
			ctx.Header("password", "admin")
			ctx.Next()
		}
	}
}
