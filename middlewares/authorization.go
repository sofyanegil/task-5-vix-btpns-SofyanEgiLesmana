package middlewares

import (
	"fmt"
	"net/http"
	"task-5-vix-btpns-SofyanEgiLesmana/app"
	"task-5-vix-btpns-SofyanEgiLesmana/database"
	"task-5-vix-btpns-SofyanEgiLesmana/models"

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

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		photoIDParam := c.Param("photoID")

		userData := c.MustGet("userData").(jwt.MapClaims)
		userIDToken, ok := userData["id"].(string)

		Photo := models.Photo{}
		fmt.Println(photoIDParam)

		err := db.Where("id = ?", photoIDParam).First(&Photo).Error

		if err != nil {
			app.ErrorResponse(c, http.StatusNotFound, "photo not found", "Not Found")
			c.Abort()
			return
		}

		if !ok || Photo.UserID != userIDToken {
			app.ErrorResponse(c, http.StatusForbidden, "you are not allowed to access this resource", "Forbidden")
			c.Abort()
			return
		}
		c.Next()
	}
}
