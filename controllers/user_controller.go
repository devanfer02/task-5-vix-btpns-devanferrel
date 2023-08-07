package controllers

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/devanfer02/task-5-vix-btpns-devanferrel/database"
	"github.com/devanfer02/task-5-vix-btpns-devanferrel/helpers"
	"github.com/devanfer02/task-5-vix-btpns-devanferrel/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUsers(ctx *gin.Context) {
	var users []models.User 

	database.DB.Find(&users)

	ctx.JSON(http.StatusOK, gin.H {
		"data": users,
		"message": "successfully fetch users",
		"status": http.StatusOK,
	})	
}

func GetUserById(ctx *gin.Context) {
	var user models.User 

	id, err := helpers.GetParamId(ctx);

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad param request",
			"status": http.StatusBadRequest,
		})
		return 
	}	

	if err := database.DB.First(&user, id).Error; err != nil {
		switch err {
			case gorm.ErrRecordNotFound :
				ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H {
					"message": "photo not found",
					"status": http.StatusNotFound,
				})
			default :
				ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H {
					"message": "internal server error",
					"status": http.StatusInternalServerError,
					"error": err.Error(),
				})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H {
		"data": user, 
		"message": "successfully fetch user",
		"status": http.StatusOK,
	})
}

func UpdateUser(ctx *gin.Context) {
	userId, ok := helpers.GetUserID(ctx);

	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
			"message": "bad body request",
			"status": http.StatusBadRequest,
		})
		return 
	}

	if _, err := govalidator.ValidateStruct(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
			"message": "bad body request", 
			"status": http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	user.Password, _= helpers.HashPassword(user.Password)

	if database.DB.Model(&user).Where("id = ?", userId).Updates(&user).RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
			"message": "failed to update data", 
			"status": http.StatusBadRequest,
		})
		return
	}

	ctx.Set("user",nil)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully update data",
		"status": http.StatusOK,
	})
}

func DeleteUser(ctx *gin.Context) {
	userId, ok := helpers.GetUserID(ctx)

	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var user models.User 

	if database.DB.Delete(&user, userId).RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
			"message": "failed to delete data",
			"status": http.StatusBadRequest,
		})
		return 
	}

	ctx.Set("user", nil)

	ctx.JSON(200, gin.H {
		"message": "successfully delete data",
		"status": http.StatusOK,
	})
}