package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PATCH opening",
	})
}
