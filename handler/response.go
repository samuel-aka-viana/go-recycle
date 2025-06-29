package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "text/json; charset=utf-8")
	ctx.JSON(code, gin.H{
		"code":      code,
		"errorCode": msg,
	})
}

func SendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "text/json; charset=utf-8")
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": fmt.Sprintf("operation from handler %s succes", op),
		"data":    data,
	})
}
