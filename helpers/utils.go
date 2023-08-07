package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetParamId(ctx *gin.Context) (int64, error) {
	idParam := ctx.Param("id")

	if id, err := strconv.ParseInt(idParam, 10, 64); err != nil {
		return 0, err;
	} else {
		return id, nil;
	}
}