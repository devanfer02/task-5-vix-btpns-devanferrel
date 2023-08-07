package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/devanfer02/task-5-vix-btpns-devanferrel/controllers"
	"github.com/devanfer02/task-5-vix-btpns-devanferrel/database"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	database.ConnectToDB();
	
	router.GET("/users/login", controllers.LoginUser)
	router.POST("/users/register", controllers.RegisterUser)
	router.PUT("/users/:userId", controllers.UpdateUser)
	router.DELETE("/users/:userId", controllers.LoginUser)

	router.GET("/photos", controllers.GetPhotos)
	router.GET("/photos/:id", controllers.GetPhotoById)
	router.POST("/photos", controllers.CreatePhoto)
	router.PUT("/photos/:id", controllers.UpdatePhoto)
	router.DELETE("/photos/:id", controllers.DeletePhoto)

	return router
}