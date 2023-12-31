package controllers

import "github.com/gin-gonic/gin"

type IUserControllers interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type IPhotoControllers interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

var (
	appJSON = "application/json"
)
