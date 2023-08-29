package controllers

import (
	"net/http"
	"task-5-vix-btpns-SofyanEgiLesmana/app"
	"task-5-vix-btpns-SofyanEgiLesmana/helpers"
	"task-5-vix-btpns-SofyanEgiLesmana/models"
	"task-5-vix-btpns-SofyanEgiLesmana/usecase"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	UserUsecase usecase.IUserUseCase
	DB          *gorm.DB
}

func NewUserController(userUsecase usecase.IUserUseCase, db *gorm.DB) *UserController {
	return &UserController{
		UserUsecase: userUsecase,
		DB:          db,
	}
}

func (controller UserController) Register(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	userReturn, err := controller.UserUsecase.Register(&User)
	if err != nil {
		app.ErrorResponse(c, http.StatusBadRequest, err.Error(), "")
		return
	}

	userRegister := map[string]interface{}{
		"id":       userReturn.ID,
		"username": userReturn.Username,
		"email":    userReturn.Email,
	}

	app.SuccessResponse(c, http.StatusCreated, "user successfully registered", userRegister)
}

func (controller UserController) Login(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	token, err := controller.UserUsecase.Login(User.Email, User.Password)

	if err != nil {
		app.ErrorResponse(c, http.StatusBadRequest, err.Error(), "")
		return
	}

	userLogin := map[string]interface{}{
		"token": token,
	}

	app.SuccessResponse(c, http.StatusOK, "login success", userLogin)
}

func (controller UserController) Update(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	userID := c.Param("userID")
	userData := c.MustGet("userData").(jwt.MapClaims)
	userIDToken := userData["id"].(string)

	var userToUpdate models.User
	if contentType == appJSON {
		c.ShouldBindJSON(&userToUpdate)
	} else {
		c.ShouldBind(&userToUpdate)
	}

	userReturn, err := controller.UserUsecase.Update(&userToUpdate, userID)
	if err != nil {
		app.ErrorResponse(c, http.StatusBadRequest, err.Error(), "")
		return
	}

	userUpdate := map[string]interface{}{
		"id":       userIDToken,
		"username": userReturn.Username,
		"email":    userReturn.Email,
	}

	app.SuccessResponse(c, http.StatusOK, "user successfully updated", userUpdate)
}

func (controller UserController) Delete(c *gin.Context) {
	userID := c.Param("userID")
	err := controller.UserUsecase.Delete(userID)
	if err != nil {
		app.ErrorResponse(c, http.StatusBadRequest, err.Error(), "")
		return
	}

	app.SuccessResponse(c, http.StatusOK, "user successfully deleted", nil)
}
