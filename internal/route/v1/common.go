package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sendSuccessResponse(ctx *gin.Context, o any) {
	ctx.JSON(http.StatusOK, o)
}

func sendFailResponse(ctx *gin.Context, e error) {
	ctx.JSON(http.StatusInternalServerError, e.Error())
}
