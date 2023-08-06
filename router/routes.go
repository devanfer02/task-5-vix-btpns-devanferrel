package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/devanfer02/task-5-vix-btpns-devanferrel/controllers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/users/login", controllers.LoginUser)
	router.POST("/users/register", controllers.RegisterUser)
	router.PUT("/users/:userId", controllers.UpdateUser)
	router.DELETE("/users/:userId", controllers.LoginUser)

	router.GET("/photos", controllers.GetPhotos)
	router.POST("photos", controllers.AddPhoto)
	router.PUT("/photos/:photoId", controllers.UpdatePhoto)
	router.DELETE("/photos/:photoId", controllers.DeletePhoto)

	return router
}