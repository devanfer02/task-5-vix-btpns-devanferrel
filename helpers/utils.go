package helpers

import (
	"strconv"

	"github.com/devanfer02/task-5-vix-btpns-devanferrel/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashed), err
}

func CompPassword(db, body *string) error {
	err := bcrypt.CompareHashAndPassword([]byte(*db), []byte(*body))

	return err
}

func GetUserID(ctx *gin.Context) (uint, bool) {
	user, _ := ctx.Get("user")
	userId := uint(0)

	switch u := user.(type) {
		case models.User :
			userId = u.ID
		default :
			return 0, false;
	}

	return userId, true
}

func GetParamId(ctx *gin.Context) (int64, error) {
	idParam := ctx.Param("id")

	if id, err := strconv.ParseInt(idParam, 10, 64); err != nil {
		return 0, err;
	} else {
		return id, nil;
	}
}