package controllers

import (
	"github.com/gin-gonic/gin"
)

func AddPhoto(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "OK",
	})
}

func GetPhotos(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "OK",
	})
}

func UpdatePhoto(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "OK",
	})
}

func DeletePhoto(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "OK",
	})
}