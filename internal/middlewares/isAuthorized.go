package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sandbox-science/deep-focus/internal/utils"
)

/*
AuthMiddleware checks if the user is
authorized to access the protected routes.
*/
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := utils.ValidateToken(c)

		if err != nil {
			c.HTML(http.StatusUnauthorized, "404.html", gin.H{
				"message": "Unauthorized: Authentication required",
			})
			fmt.Println(err)
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
