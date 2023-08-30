package router

import (
	"task-5-vix-btpns-SofyanEgiLesmana/controllers"
	"task-5-vix-btpns-SofyanEgiLesmana/database"
	"task-5-vix-btpns-SofyanEgiLesmana/middlewares"
	"task-5-vix-btpns-SofyanEgiLesmana/repository"
	"task-5-vix-btpns-SofyanEgiLesmana/usecase"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	db := database.GetDB()

	userRepository := repository.NewUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepository, db)
	userController := controllers.NewUserController(userUsecase, db)

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", userController.Register)
		userRouter.POST("/login", userController.Login)
		userRouter.Use(middlewares.Authentication())
		userRouter.PUT("/:userID", middlewares.UserAuthorization(), userController.Update)
		userRouter.DELETE("/:userID", middlewares.UserAuthorization(), userController.Delete)
	}

	photoRepository := repository.NewPhotoRepository()
	photoUsecase := usecase.NewPhotoUsecase(photoRepository, db)
	photoController := controllers.NewPhotoController(photoUsecase, db)

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", photoController.Create)
		photoRouter.GET("/", photoController.GetAll)
		photoRouter.PUT("/:photoID", middlewares.PhotoAuthorization(), photoController.Update)
		photoRouter.DELETE("/:photoID", middlewares.PhotoAuthorization(), photoController.Delete)
	}

	return r
}
