package controllers

import (
	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H {
		"message": "OK",
	})
}

func LoginUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "OK",
	})
}

func UpdateUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "OK",
	})
}

func DeleteUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H {
		"message":"OK",
	})
}