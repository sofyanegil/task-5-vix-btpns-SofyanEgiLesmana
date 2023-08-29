package middlewares

import (
	"net/http"
	"task-5-vix-btpns-SofyanEgiLesmana/app"
	"task-5-vix-btpns-SofyanEgiLesmana/helpers"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			app.ErrorResponse(c, http.StatusUnauthorized, err.Error(), "Unauthorized")
			return
		}
		c.Set("userData", verifyToken)
		c.Next()
	}
}
