package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BulkUpdateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PUT opening",
	})
}
