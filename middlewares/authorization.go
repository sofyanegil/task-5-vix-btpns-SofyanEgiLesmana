package middlewares

import (
	"net/http"
	"task-5-vix-btpns-SofyanEgiLesmana/app"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDParam := c.Param("userID")
		userData := c.MustGet("userData").(jwt.MapClaims)
		userIDToken, ok := userData["id"].(string)
		if !ok || userIDToken != userIDParam {
			app.ErrorResponse(c, http.StatusForbidden, "you are not allowed to access this resource", "Forbidden")
			c.Abort()
			return
		}
		c.Next()
	}
}
