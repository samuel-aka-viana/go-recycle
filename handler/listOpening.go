package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOpeningsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "openings",
	})
}
