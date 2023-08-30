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

type PhotoController struct {
	PhotoUsecase usecase.IPhotoUseCase
	DB           *gorm.DB
}

func NewPhotoController(photoUsecase usecase.IPhotoUseCase, db *gorm.DB) *PhotoController {
	return &PhotoController{
		PhotoUsecase: photoUsecase,
		DB:           db,
	}
}

func (controller PhotoController) Create(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := userData["id"].(string)

	photoReturn, err := controller.PhotoUsecase.Create(&Photo, userID)

	if err != nil {
		app.ErrorResponse(c, http.StatusBadRequest, err.Error(), "")
		return
	}

	app.SuccessResponse(c, http.StatusCreated, "photo successfully created", photoReturn)
}

func (controller PhotoController) GetAll(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	photoReturn, err := controller.PhotoUsecase.GetAll()

	if err != nil {
		app.ErrorResponse(c, http.StatusBadRequest, err.Error(), "")
		return
	}
	app.SuccessResponse(c, http.StatusOK, "photos successfully retrieved", photoReturn)
}

func (controller PhotoController) Update(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoID := c.Param("photoID")
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := userData["id"].(string)

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = photoID

	photoReturn, err := controller.PhotoUsecase.Update(&Photo, photoID)

	if err != nil {
		app.ErrorResponse(c, http.StatusBadRequest, err.Error(), "")
		return
	}

	app.SuccessResponse(c, http.StatusOK, "photo successfully updated", photoReturn)
}

func (controller PhotoController) Delete(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoID := c.Param("photoID")
	userID := userData["id"].(string)

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = photoID

	err := controller.PhotoUsecase.Delete(photoID)

	if err != nil {
		app.ErrorResponse(c, http.StatusBadRequest, err.Error(), "")
		return
	}

	app.SuccessResponse(c, http.StatusOK, "photo successfully deleted", nil)
}
